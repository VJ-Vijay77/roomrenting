package controller

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/VJ-Vijay77/r4room/pkg/controllertwo/orderid"
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
	RID, _ := strconv.Atoi(Roomid)
	Usersid, _ := strconv.Atoi(UserID)

	TP, _ := strconv.Atoi(TotalPrice)

	//taking order details
	var roomnames string
	// db.Raw("SELECT rooms.room_name FROM carts INNER JOIN rooms ON rooms.id=carts.cartsroomid WHERE carts.user_id=?", Usersid).Scan(&roomnames)
	// var sendinginfo string
	// for i := range roomnames {
	// 	sendinginfo += roomnames[i]
	// }
	db.Raw("SELECT room_name FROM rooms WHERE id=?", RID).Scan(&roomnames)

	//parsing address of the user
	var address []models.Useraddress
	db.Where("user_id=?", UserID).Find(&address)

	//taking wallet balance
	var wallet models.Wallets
	var walletbal int
	db.Select("balance").Where("user_id=?", Usersid).Find(&wallet)

	if wallet.Balance >= TP {
		walletbal = TP - 1
	} else if wallet.Balance < TP {
		walletbal = wallet.Balance
	}

	//getting category of room
	var category string
	db.Raw("SELECT category FROM rooms WHERE id=?", RID).Scan(&category)

	//gettting coupons
	var coupons []models.Coupons
	db.Find(&coupons)

	c.HTML(200, "payment.gohtml", gin.H{
		"total":     TotalPrice,
		"roomnames": roomnames,
		"allrooms":  roomnames,
		"username":  UserName,
		"address":   address,
		"uid":       IDofUser,
		"startdate": Startdate,
		"endate":    Endate,
		"roomid":    Roomid,
		"wbal":      walletbal,
		"coupon":    coupons,
		"category":  category,
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

	// var roomsids []models.Carts
	// db.Select("cartsroomid").Where("user_id=?", UserID).Find(&roomsids)

	start := c.Param("start")
	end := c.Param("end")
	wallet := c.Param("wallet")
	wtotal := c.Param("total")

	var Wallbal int
	var zero int = 0
	var walbalets models.Wallets
	if wallet == "0" {
		Wallbal = 0
	} else {
		k, _ := strconv.Atoi(wallet)
		Wallbal = k
		db.Raw("UPDATE wallets SET balance=? WHERE user_id=?", zero, UserID).Scan(&walbalets)
	}

	//fmt.Println(start,end)
	FName := c.PostForm("firstnamef")
	LName := c.PostForm("lastnamef")
	HName := c.PostForm("housenamef")
	Place := c.PostForm("placef")
	State := c.PostForm("statef")
	Mobile := c.PostForm("mobilef")
	Roomid := c.PostForm("roomid")
	RID, _ := strconv.Atoi(Roomid)
	Allrooms := c.PostForm("roomnames")
	total, _ := strconv.Atoi(wtotal)

	rand.Seed(time.Now().UnixNano())
	Refid := orderid.OrderIdGeneration(5)
	RefereceId := "PAO_" + Refid

	var status = "checkedin"

	var idsofroom models.Orderedrooms

	idsofroom.Roomid = RID
	idsofroom.User_Id = UserID
	idsofroom.Status = status
	idsofroom.Checkindate = start
	idsofroom.Checkoutdate = end
	db.Select("roomid", "user_id", "status", "checkindate", "checkoutdate").Create(&idsofroom)

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
	orderdetails.Roomid = RID
	orderdetails.Refid = RefereceId

	db.Select("user_id", "firstname", "lastname", "housename", "place", "state", "mobile", "totalprice", "paymentmethod", "roomnames", "accountholder", "checkindate", "checkoutdate", "wallet", "roomid", "refid").Create(&orderdetails)

	var checkoutdate models.Rooms
	db.Raw("UPDATE rooms SET checkoutdate=? WHERE id=?", end, RID).Scan(&checkoutdate)

	var deletefromcart models.Carts
	db.Raw("DELETE FROM carts WHERE cartsroomid =?", RID).Scan(&deletefromcart)

	var updateroomstatus models.Rooms

	db.Raw("UPDATE rooms SET status='booked' WHERE id=?", RID).Scan(&updateroomstatus)

	ok := RefereceId
	k, _ := json.Marshal(ok)
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.Write(k)

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

	RefId := c.Param("refid")

	var details models.Orders
	db.Where("refid=?", RefId).Find(&details)

	c.HTML(200, "paymentsuccess.gohtml", gin.H{
		"username": UserName,
		"count":    count,
		"wcount":   wishlistcount,
		"details":  details,
		"email":    userID,
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
	RID, _ := strconv.Atoi(Roomid)

	UID, _ := strconv.Atoi(UserID)

	//taking order details
	var roomnames string
	// db.Raw("SELECT rooms.room_name FROM carts INNER JOIN rooms ON rooms.id=carts.cartsroomid WHERE carts.user_id=?", UID).Scan(&roomnames)
	// var sendinginfo string
	// for i := range roomnames {
	// 	sendinginfo += roomnames[i]
	// }
	db.Raw("SELECT room_name FROM rooms WHERE id=?", RID).Scan(&roomnames)

	//parsing address of the user
	var address []models.Useraddress
	db.Where("user_id=?", UID).Find(&address)

	addressID := c.Request.FormValue("radioaddress")
	AID, _ := strconv.Atoi(addressID)
	var addresspick models.Useraddress
	db.Where("adrid=?", AID).Find(&addresspick)

	//taking wallet balance
	var wallet models.Wallets
	db.Select("balance").Where("user_id=?", UID).Find(&wallet)

	//getting category of room
	var category string
	db.Raw("SELECT category FROM rooms WHERE id=?", RID).Scan(&category)

	//gettting coupons
	var coupons []models.Coupons
	db.Find(&coupons)

	c.HTML(200, "paymentaddressfill.gohtml", gin.H{
		"total":     TotalPrice,
		"roomnames": roomnames,
		"allrooms":  roomnames,
		"username":  UserName,
		"lastname":  LastName,
		"address":   address,
		"adr":       addresspick,
		"uid":       IDofuser,
		"startdate": Startdate,
		"endate":    Endate,
		"roomid":    Roomid,
		"wbal":      wallet.Balance,
		"coupon":    coupons,
		"category":  category,
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
	Roomid := c.Param("roomid")
	RID, _ := strconv.Atoi(Roomid)
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
		"start":    Startdate,
		"end":      Endate,
		"wallet":   Wallet,
		"roomid":   RID,
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
	Roomid := c.Param("roomid")
	RID, _ := strconv.Atoi(Roomid)
	var Wallbal int
	var zero int = 0
	var walbalets models.Wallets
	if Wallet == "0" {
		Wallbal = 0
	} else {
		k, _ := strconv.Atoi(Wallet)
		Wallbal = k
		db.Raw("UPDATE wallets SET balance=? WHERE user_id=?", zero, UserID).Scan(&walbalets)
	}

	totalint, _ := strconv.Atoi(Total)
	GrandTotal := totalint / 100

	rand.Seed(time.Now().UnixNano())
	Refid := orderid.OrderIdGeneration(5)
	RefereceId := "RP_" + Refid

	var status = "checkedin"

	var idsofroom models.Orderedrooms

	idsofroom.Roomid = RID
	idsofroom.User_Id = UserID
	idsofroom.Status = status
	idsofroom.Checkindate = Startdate
	idsofroom.Checkoutdate = Endate
	db.Select("roomid", "user_id", "status", "checkindate", "checkoutdate").Create(&idsofroom)

	var roomnames string
	db.Raw("SELECT room_name FROM rooms WHERE id=?", RID).Scan(&roomnames)

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
	orderdetails.Roomnames = roomnames
	orderdetails.Accountholder = username
	orderdetails.Checkindate = Startdate
	orderdetails.Checkoutdate = Endate
	orderdetails.Wallet = Wallbal
	orderdetails.Roomid = RID
	orderdetails.Refid = RefereceId

	db.Select("user_id", "firstname", "lastname", "housename", "place", "state", "mobile", "totalprice", "paymentmethod", "roomnames", "accountholder", "checkindate", "checkoutdate", "wallet", "roomid", "refid").Create(&orderdetails)

	var checkoutdate models.Rooms
	db.Raw("UPDATE rooms SET checkoutdate=? WHERE id=?", Endate, RID).Scan(&checkoutdate)

	var deletefromcart models.Carts
	db.Raw("DELETE FROM carts WHERE cartsroomid =?", RID).Scan(&deletefromcart)

	var updateroomstatus models.Rooms

	db.Raw("UPDATE rooms SET status='booked' WHERE id=?", RID).Scan(&updateroomstatus)

	db.AutoMigrate(models.Razorpaydetails{})
	var razor models.Razorpaydetails
	razor.User_Id = UserID
	razor.Payid = PayId
	razor.Orderid = OrderId
	razor.Signature = Signature
	db.Select("user_id", "payid", "orderid", "signature").Create(&razor)

	c.Redirect(303, "/user/payment/razorpay/successpage/"+PayId+"/"+OrderId+"/"+Signature+"/"+RefereceId)

}

func RPSuccess(c *gin.Context) {

	db := database.GetDb()
	session, _ := Store.Get(c.Request, "session")
	useriD := session.Values["userID"]
	userID := fmt.Sprintf("%s", useriD)
	var userinfos models.Users

	db.Raw("SELECT first_name,last_name FROM users where email=?", userID).Scan(&userinfos)
	username := userinfos.First_Name

	var UserID int
	db.Raw("SELECT id FROM users WHERE email=?", userID).Scan(&UserID)

	//cart count
	var count int
	db.Raw("SELECT COUNT(user_id) FROM carts WHERE user_id=?", UserID).Scan(&count)
	//wishlist count
	var wishlistcount int
	db.Raw("SELECT COUNT(user_id) FROM wishlists WHERE user_id=?", UserID).Scan(&wishlistcount)

	PID := c.Param("pid")
	OrdID := c.Param("orderid")
	Sign := c.Param("signature")
	Refid := c.Param("refid")

	var orderdetails models.Orders
	db.Where("refid=?", Refid).Find(&orderdetails)

	c.HTML(200, "successpage.gohtml", gin.H{
		"payid":    PID,
		"ordid":    OrdID,
		"sign":     Sign,
		"refid":    Refid,
		"username": username,
		"count":    count,
		"wcount":   wishlistcount,
		"details":  orderdetails,
		"email":    userID,
	})
}
