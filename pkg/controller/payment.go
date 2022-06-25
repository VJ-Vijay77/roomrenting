package controller

import (
	"fmt"
	"os"
	"strconv"

	"github.com/VJ-Vijay77/r4room/pkg/database"
	"github.com/VJ-Vijay77/r4room/pkg/models"
	"github.com/gin-gonic/gin"
	razorpay "github.com/razorpay/razorpay-go"
	
)

func Payment(c *gin.Context) {
	db := database.GetDb()
	//checking session
	ok := UserLogedCheck(c)
	if !ok {
		c.HTML(200, "userhome.gohtml", nil)
		return
	}
	session, _ := Store.Get(c.Request, "session")
	useriD := session.Values["userID"]
	userID := fmt.Sprintf("%s", useriD)
	var userinfos models.Users

	db.Raw("SELECT id,first_name FROM users where email=?", userID).Scan(&userinfos)
	UserName := userinfos.First_Name
	IDofUser := userinfos.ID

	TotalPrice := c.Param("TotalPrice")
	UserID := c.Param("UserID")
	Roomid := c.Param("Roomid")
	Startdate := c.Param("Startdate")
	Endate := c.Param("Endate")
	RID,_ := strconv.Atoi(Roomid) 
	Usersid, _ := strconv.Atoi(UserID)
	//taking order details
	var roomnames []string
	db.Raw("SELECT rooms.room_name FROM carts INNER JOIN rooms ON rooms.id=carts.cartsroomid WHERE carts.user_id=?", Usersid).Scan(&roomnames)
	var sendinginfo string
	for i := range roomnames {
		sendinginfo += roomnames[i]
	}


	//parsing address of the user
	var address []models.Useraddress
	db.Where("user_id=?", UserID).Find(&address)

	//taking wallet balance
	var wallet models.Wallets
	db.Select("balance").Where("user_id=?",Usersid).Find(&wallet)

	//getting category of room
	var category string
	db.Raw("SELECT category FROM rooms WHERE id=?",RID).Scan(&category)

	//gettting coupons
	var coupons []models.Coupons
	db.Find(&coupons)

	c.HTML(200, "payment.gohtml", gin.H{
		"total":     TotalPrice,
		"roomnames": roomnames,
		"allrooms":  sendinginfo,
		"username":  UserName,
		"address":   address,
		"uid":       IDofUser,
		"startdate": Startdate,
		"endate":    Endate,
		"roomid":    Roomid,
		"wbal":wallet.Balance,
		"coupon":coupons,
		"category":category,
	})
}



func PaymentConfirm(c *gin.Context) {
	db := database.GetDb()
	session, _ := Store.Get(c.Request, "session")
	useriD := session.Values["userID"]
	userID := fmt.Sprintf("%s", useriD)
	var userinfos models.Users

	db.Raw("SELECT first_name FROM users where email=?", userID).Scan(&userinfos)
	UserName := userinfos.First_Name

	var UserID int
	db.Raw("SELECT id FROM users WHERE email=?", userID).Scan(&UserID)

	var roomsids []models.Carts
	db.Select("cartsroomid").Where("user_id=?", UserID).Find(&roomsids)

	

	
	start := c.Param("start")
	end := c.Param("end")
	wallet := c.Param("wallet")
	wtotal := c.Param("total")
	
	var Wallbal int 
	var zero int= 0
	var walbalets models.Wallets
	if wallet == "0"{
		Wallbal = 0
	}else{
		k,_ := strconv.Atoi(wallet)
		Wallbal = k
		db.Raw("UPDATE wallets SET balance=? WHERE user_id=?",zero,UserID).Scan(&walbalets)
	}



	//fmt.Println(start,end)
	 FName := c.PostForm("firstnamef")	
	LName := c.PostForm("lastnamef")
	HName := c.PostForm("housenamef")
	Place := c.PostForm("placef")
	State := c.PostForm("statef")
	Mobile := c.PostForm("mobilef")
	// Total := c.PostForm("totalprice")
	Allrooms := c.PostForm("allroomnames")
	total,_ := strconv.Atoi(wtotal) 

	var status = "checkedin"

	var idsofroom models.Orderedrooms
	for _, val := range roomsids {
		idsofroom.Roomid = val.Cartsroomid
		idsofroom.User_Id = UserID
		idsofroom.Status = status
		idsofroom.Checkindate = start
		idsofroom.Checkoutdate = end
		db.Select("roomid", "user_id", "status","checkindate","checkoutdate").Create(&idsofroom)
	}
	

	var paymode string = "Payment At Office"

	var orderdetails models.Orders
	orderdetails.User_ID = UserID
	orderdetails.Firstname = FName
	orderdetails.Lastname = LName
	orderdetails.Housename = HName
	orderdetails.Place = Place
	orderdetails.State = State
	orderdetails.Mobile = Mobile
	orderdetails.Totalprice = total
	orderdetails.Paymentmethod = paymode
	orderdetails.Roomnames = Allrooms
	orderdetails.Accountholder = UserName
	orderdetails.Checkindate = start 
	orderdetails.Checkoutdate = end 
	orderdetails.Wallet = Wallbal
	db.Select("user_id", "firstname", "lastname", "housename", "place", "state", "mobile", "totalprice","paymentmethod","roomnames", "accountholder","checkindate","checkoutdate","wallet").Create(&orderdetails)

	

	var checkoutdate models.Rooms
	db.Raw("UPDATE rooms SET checkoutdate=? WHERE room_name=?",end,Allrooms).Scan(&checkoutdate)

	var deletefromcart models.Carts
	db.Raw("DELETE FROM carts WHERE user_id=?", UserID).Scan(&deletefromcart)

	var updateroomstatus models.Rooms

	for _, val := range roomsids {

		db.Raw("UPDATE rooms SET status='booked' WHERE id=?", val.Cartsroomid).Scan(&updateroomstatus)
	}

}

func PaymentSuccess(c *gin.Context) {
	db := database.GetDb()
	session, _ := Store.Get(c.Request, "session")
	useriD := session.Values["userID"]
	userID := fmt.Sprintf("%s", useriD)
	var userinfos models.Users

	db.Raw("SELECT first_name FROM users where email=?", userID).Scan(&userinfos)
	UserName := userinfos.First_Name

	var UserID int
	db.Raw("SELECT id FROM users WHERE email=?", userID).Scan(&UserID)

	//cart count
	var count int
	db.Raw("SELECT COUNT(user_id) FROM carts WHERE user_id=?", UserID).Scan(&count)
	//wishlist count
	var wishlistcount int
	db.Raw("SELECT COUNT(user_id) FROM wishlists WHERE user_id=?", UserID).Scan(&wishlistcount)

	c.HTML(200, "paymentsuccess.gohtml", gin.H{
		"username": UserName,
		"count":    count,
		"wcount":   wishlistcount,
	})
}

func PaymentAddressPick(c *gin.Context) {
	db := database.GetDb()
	//checking session
	ok := UserLogedCheck(c)
	if !ok {
		c.HTML(200, "userhome.gohtml", nil)
		return
	}
	session, _ := Store.Get(c.Request, "session")
	useriD := session.Values["userID"]
	userID := fmt.Sprintf("%s", useriD)
	var userinfos models.Users

	db.Raw("SELECT id,first_name,last_name FROM users where email=?", userID).Scan(&userinfos)
	UserName := userinfos.First_Name
	LastName := userinfos.Last_Name
	IDofuser := userinfos.ID

	TotalPrice := c.Param("TotalPrice")
	UserID := c.Param("UserID")
	Roomid := c.Param("Roomid")
	Startdate := c.Param("Startdate")
	Endate := c.Param("Endate")
	//fmt.Println("hiii")
	fmt.Println(TotalPrice,UserID,Roomid,Startdate,Endate)
	UID, _ := strconv.Atoi(UserID)

	//taking order details
	var roomnames []string
	db.Raw("SELECT rooms.room_name FROM carts INNER JOIN rooms ON rooms.id=carts.cartsroomid WHERE carts.user_id=?", UID).Scan(&roomnames)
	var sendinginfo string
	for i := range roomnames {
		sendinginfo += roomnames[i]
	}

	//parsing address of the user
	var address []models.Useraddress
	db.Where("user_id=?", UID).Find(&address)

	addressID := c.Request.FormValue("radioaddress")
	AID, _ := strconv.Atoi(addressID)
	var addresspick models.Useraddress
	db.Where("adrid=?", AID).Find(&addresspick)

	//taking wallet balance
	var wallet models.Wallets
	db.Select("balance").Where("user_id=?",UID).Find(&wallet)

	c.HTML(200, "paymentaddressfill.gohtml", gin.H{
		"total":     TotalPrice,
		"roomnames": roomnames,
		"allrooms":  sendinginfo,
		"username":  UserName,
		"lastname":  LastName,
		"address":   address,
		"adr":       addresspick,
		"uid":       IDofuser,
		"startdate": Startdate,
		"endate":    Endate,
		"roomid":    Roomid,
		"wbal":wallet.Balance,
	})
}

func RazorPay(c *gin.Context) {

	db := database.GetDb()

	type PageVariables struct {
		OrderId string
	}
	Total := c.Param("total")
	Userid := c.Param("uid")
	Startdate := c.Param("start")
	Endate := c.Param("end")
	UID, _ := strconv.Atoi(Userid)
	Grandtotal := Total + "00"
	total, _ := strconv.Atoi(Grandtotal)
	Wallet := c.Param("wallet")

	var userinfos models.Users
	db.Where("id=?", UID).Find(&userinfos)

	var address models.Useraddress
	db.Where("user_id=? AND main='true'", UID).Find(&address)

	client := razorpay.NewClient("rzp_test_pOTtQfNqdwqxCR", "PShKIEvMv6DQsRQSCpAxWjXX")

	datas := map[string]interface{}{
		"amount":   total,
		"currency": "INR",
		"receipt":  "receipt_id",
	}
	//   body,_ := client.Order.Create(data)

	body, err := client.Order.Create(datas, nil)
	fmt.Println(body)

	if err != nil {
		fmt.Printf("problem %v \n", err)
		os.Exit(1)
	}

	value := body["id"]

	str := value.(string)

	HomePageVars := PageVariables{
		OrderId: str,
	}

	c.HTML(200, "razorpay.gohtml", gin.H{
		"home":     HomePageVars,
		"total":    total,
		"username": userinfos.First_Name,
		"email":    userinfos.Email,
		"mobile":   userinfos.Mobile,
		"address":  address,
		"start":Startdate,
		"end":Endate,
		"wallet":Wallet,
	})

}

func RazorPaySuccess(c *gin.Context) {
	db := database.GetDb()
	session, _ := Store.Get(c.Request, "session")
	useriD := session.Values["userID"]
	userID := fmt.Sprintf("%s", useriD)
	var userinfos models.Users

	db.Raw("SELECT first_name,last_name FROM users where email=?", userID).Scan(&userinfos)
	username := userinfos.First_Name
	lastname := userinfos.Last_Name

	var UserID int
	db.Raw("SELECT id FROM users WHERE email=?", userID).Scan(&UserID)


	PayId := c.Param("rpid")
	OrderId := c.Param("roid")
	Signature := c.Param("rsign")
	Startdate := c.Param("start")
	Endate := c.Param("end")
	Total := c.Param("total")
	Wallet := c.Param("wallet")
	var Wallbal int 
	var zero int= 0
	var walbalets models.Wallets
	if Wallet == "0"{
		Wallbal = 0
	}else{
		k,_ := strconv.Atoi(Wallet)
		Wallbal = k
		db.Raw("UPDATE wallets SET balance=? WHERE user_id=?",zero,UserID).Scan(&walbalets)
	}


	totalint,_ := strconv.Atoi(Total)
	GrandTotal := totalint/100

	

	
	
	var roomsids []models.Carts
	db.Select("cartsroomid").Where("user_id=?", UserID).Find(&roomsids)

	var status = "checkedin"

	var idsofroom models.Orderedrooms
	for _, val := range roomsids {
		idsofroom.Roomid = val.Cartsroomid
		idsofroom.User_Id = UserID
		idsofroom.Status = status
		idsofroom.Checkindate = Startdate
		idsofroom.Checkoutdate = Endate
		db.Select("roomid", "user_id", "status","checkindate","checkoutdate").Create(&idsofroom)
	}

	var roomnames []string
	db.Raw("SELECT rooms.room_name FROM carts INNER JOIN rooms ON rooms.id=carts.cartsroomid WHERE carts.user_id=?", UserID).Scan(&roomnames)
	var sendinginfo string
	for i := range roomnames {
		sendinginfo += roomnames[i]
	}
	

		var paymode string = "Razor Pay"
	var orderdetails models.Orders
	var adress models.Useraddress
	db.Where("user_id=? AND main='true'", UserID).Find(&adress)
	orderdetails.User_ID = UserID
	orderdetails.Firstname = username
	orderdetails.Lastname = lastname
	orderdetails.Housename = adress.Housename
	orderdetails.Place = adress.Place
	orderdetails.State = adress.State
	orderdetails.Mobile = adress.Mobile
	orderdetails.Totalprice = GrandTotal
	orderdetails.Paymentmethod = paymode
	orderdetails.Roomnames = sendinginfo
	orderdetails.Accountholder = username
	orderdetails.Checkindate = Startdate
	orderdetails.Checkoutdate = Endate
	orderdetails.Wallet = Wallbal

	db.Select("user_id", "firstname", "lastname", "housename", "place", "state", "mobile", "totalprice","paymentmethod","roomnames", "accountholder","checkindate","checkoutdate","wallet").Create(&orderdetails)

	var checkoutdate models.Rooms
	db.Raw("UPDATE rooms SET checkoutdate=? WHERE room_name=?",Endate,sendinginfo).Scan(&checkoutdate)


	var deletefromcart models.Carts
	db.Raw("DELETE FROM carts WHERE user_id=?", UserID).Scan(&deletefromcart)

	var updateroomstatus models.Rooms

	for _, val := range roomsids {

		db.Raw("UPDATE rooms SET status='booked' WHERE id=?", val.Cartsroomid).Scan(&updateroomstatus)
	}

	db.AutoMigrate(models.Razorpaydetails{})
	var razor models.Razorpaydetails
	razor.User_Id = UserID
	razor.Payid = PayId
	razor.Orderid = OrderId
	razor.Signature = Signature
	db.Select("user_id", "payid", "orderid", "signature").Create(&razor)

	
	c.Redirect(303, "/user/payment/success")

}
