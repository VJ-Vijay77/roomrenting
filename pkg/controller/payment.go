package controller

import (
	"fmt"

	"github.com/VJ-Vijay77/r4room/pkg/database"
	"github.com/VJ-Vijay77/r4room/pkg/models"
	"github.com/gin-gonic/gin"
)

func Payment(c *gin.Context) {
	//checking session
	ok := UserLogedCheck(c)
	if !ok {
		c.HTML(200, "userhome.gohtml", nil)
		return
	}

	TotalPrice := c.Param("TotalPrice")
	UserID := c.Param("UserID")

	db := database.GetDb()
	//taking order details
	var roomnames []string
	db.Raw("SELECT rooms.room_name FROM carts INNER JOIN rooms ON rooms.id=carts.cartsroomid WHERE carts.user_id=?", UserID).Scan(&roomnames)
	var sendinginfo string
	for i:= range roomnames {
		sendinginfo += roomnames[i]
	}


	c.HTML(200, "payment.gohtml", gin.H{
		"total":     TotalPrice,
		"roomnames": roomnames,
		"allrooms":sendinginfo,
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
	db.Select("cartsroomid").Where("user_id=?",UserID).Find(&roomsids)
	
	
	var idsofroom models.Orderedrooms
	for _,val := range roomsids{
		idsofroom.Roomid=val.Cartsroomid
		idsofroom.User_Id=UserID
		db.Select("roomid","user_id").Create(&idsofroom)
	}
	

	FName := c.PostForm("firstnamef")
	LName := c.PostForm("lastnamef")
	HName := c.PostForm("housenamef")
	Place := c.PostForm("placef")
	State := c.PostForm("statef")
	Mobile := c.PostForm("mobilef")
	Total := c.PostForm("totalprice")
	Allrooms := c.PostForm("allroomnames")

	var orderdetails models.Orders
	orderdetails.User_ID=UserID
	orderdetails.Firstname=FName
	orderdetails.Lastname=LName
	orderdetails.Housename=HName
	orderdetails.Place=Place
	orderdetails.State=State
	orderdetails.Mobile=Mobile
	orderdetails.Totalprice=Total
	orderdetails.Roomnames=Allrooms
	orderdetails.Accountholder=UserName
	 db.Select("user_id","firstname","lastname","housename","place","state","mobile","totalprice","roomnames","accountholder").Create(&orderdetails)

	 var deletefromcart models.Carts
	db.Raw("DELETE FROM carts WHERE user_id=?",UserID).Scan(&deletefromcart)
	
	var updateroomstatus models.Rooms 
	
	for _,val := range roomsids{
		
		db.Raw("UPDATE rooms SET status='booked' WHERE id=?",val.Cartsroomid).Scan(&updateroomstatus)
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
	db.Raw("SELECT COUNT(user_id) FROM carts WHERE user_id=?",UserID).Scan(&count)
	//wishlist count
	var wishlistcount int
	db.Raw("SELECT COUNT(user_id) FROM wishlists WHERE user_id=?",UserID).Scan(&wishlistcount)
	


	c.HTML(200,"paymentsuccess.gohtml",gin.H{
		"username":UserName,
		"count":count,
		"wcount":wishlistcount,
	})
}