package controller

import (
	"fmt"
	"log"
	"time"

	"github.com/VJ-Vijay77/r4room/pkg/database"
	"github.com/VJ-Vijay77/r4room/pkg/models"
	"github.com/gin-gonic/gin"
)

func AvailableSearchRooms(c *gin.Context) {
	db := database.GetDb()

	session, err := Store.Get(c.Request, "session")
	if err != nil {
		log.Println("Cannot get sessions!")
	}
	useriD := session.Values["userID"]
	userID := fmt.Sprintf("%s", useriD)
	var userinfos models.Users

	db.Raw("SELECT first_name FROM users where email=?", userID).Scan(&userinfos)

	UserName := userinfos.First_Name
	var rooms []models.Rooms
	var status = "available"
	db.Where("status=?", status).Find(&rooms)

	var UserID int
	db.Raw("SELECT id FROM users WHERE email=?", userID).Scan(&UserID)
	//cart count
	var count int
	db.Raw("SELECT COUNT(user_id) FROM carts WHERE user_id=?", UserID).Scan(&count)
	//wishlist count
	var wishlistcount int
	db.Raw("SELECT COUNT(user_id) FROM wishlists WHERE user_id=?", UserID).Scan(&wishlistcount)

	currentTime := time.Now()
	cdate := currentTime.Format("2006-01-02")
	c.HTML(200, "roomsearchpage.gohtml", gin.H{
		"rooms":    rooms,
		"username": UserName,
		"count":    count,
		"wcount":   wishlistcount,
		"cdate" :cdate,
	})
}

func BookedSearchRooms(c *gin.Context) {
	db := database.GetDb()

	session, err := Store.Get(c.Request, "session")
	if err != nil {
		log.Println("Cannot get sessions!")
	}
	useriD := session.Values["userID"]
	userID := fmt.Sprintf("%s", useriD)
	var userinfos models.Users

	db.Raw("SELECT first_name FROM users where email=?", userID).Scan(&userinfos)

	UserName := userinfos.First_Name
	var rooms []models.Rooms
	var status = "booked"
	db.Where("status=?", status).Find(&rooms)

	var UserID int
	db.Raw("SELECT id FROM users WHERE email=?", userID).Scan(&UserID)
	//cart count
	var count int
	db.Raw("SELECT COUNT(user_id) FROM carts WHERE user_id=?", UserID).Scan(&count)
	//wishlist count
	var wishlistcount int
	db.Raw("SELECT COUNT(user_id) FROM wishlists WHERE user_id=?", UserID).Scan(&wishlistcount)


	currentTime := time.Now()
	cdate := currentTime.Format("2006-01-02")
	c.HTML(200, "bookedrooms.gohtml", gin.H{
		"rooms":    rooms,
		"username": UserName,
		"count":    count,
		"wcount":   wishlistcount,
		"cdate":cdate,
	})
}

func RoomInfo(c *gin.Context) {
	db := database.GetDb()
	ID := c.Param("ID")

	session, _ := Store.Get(c.Request, "session")
	useriD := session.Values["userID"]
	userID := fmt.Sprintf("%s", useriD)
	var userinfos models.Users
	db.Raw("SELECT first_name FROM users where email=?", userID).Scan(&userinfos)
	UserName := userinfos.First_Name

	var roominfo models.Rooms
	//db.Raw("SELECT room_name,id,room_price,category,cover,sub1,sub2,sub3,sub4,sub5 FROM rooms WHERE id=?", ID).Scan(&roominfo)
	db.Where("id=?", ID).Find(&roominfo)

	var UserID int
	db.Raw("SELECT id FROM users WHERE email=?", userID).Scan(&UserID)
	var count int
	db.Raw("SELECT COUNT(user_id) FROM carts WHERE user_id=?", UserID).Scan(&count)
	//wishlist count
	var wishlistcount int
	db.Raw("SELECT COUNT(user_id) FROM wishlists WHERE user_id=?", UserID).Scan(&wishlistcount)

	currentTime := time.Now()
	cdate := currentTime.Format("2006-01-02")
	cdate2 := currentTime.AddDate(0, 0, 1)
	nextday := cdate2.Format("2006-01-02")

	c.HTML(200, "roominfo.gohtml", gin.H{
		"roomname":     roominfo.Room_Name,
		"roomid":       roominfo.ID,
		"roomcategory": roominfo.Category,
		"roomprice":    roominfo.Room_Price,
		"cover":        roominfo.Cover,
		"sub1":         roominfo.Sub1,
		"sub2":         roominfo.Sub2,
		"sub3":         roominfo.Sub3,
		"sub4":         roominfo.Sub4,
		"sub5":         roominfo.Sub5,
		"status":       roominfo.Status,
		"checkout":     roominfo.Checkoutdate,
		"desc":			roominfo.Description,
		"Offers":roominfo.Offers,
		"off":roominfo.Value,
		"save":roominfo.Room_Price - roominfo.Discountprice,
		"discount":roominfo.Discountprice,
		"username":     UserName,
		"count":        count,
		"wcount":       wishlistcount,
		"cdate":        cdate,
		"ndate":        nextday,
	})
}

func AllSearchRooms(c *gin.Context) {

	db := database.GetDb()

	session, err := Store.Get(c.Request, "session")
	if err != nil {
		log.Println("Cannot get sessions!")
	}
	useriD := session.Values["userID"]
	userID := fmt.Sprintf("%s", useriD)
	var userinfos models.Users

	db.Raw("SELECT first_name FROM users where email=?", userID).Scan(&userinfos)

	UserName := userinfos.First_Name

	var allrooms []models.Rooms
	db.Find(&allrooms)

	var UserID int
	db.Raw("SELECT id FROM users WHERE email=?", userID).Scan(&UserID)
	//cart count
	var count int
	db.Raw("SELECT COUNT(user_id) FROM carts WHERE user_id=?", UserID).Scan(&count)
	//wishlist count
	var wishlistcount int
	db.Raw("SELECT COUNT(user_id) FROM wishlists WHERE user_id=?", UserID).Scan(&wishlistcount)

	currentTime := time.Now()
	cdate := currentTime.Format("2006-01-02")

	for _,i := range allrooms{
		if i.Offers == "true"{
			i.Discountprice = i.Room_Price * (100-i.Value)/100
			db.Raw("UPDATE rooms SET discountprice=? WHERE id=?",i.Discountprice,i.ID).Scan(allrooms)

		}
	}


	c.HTML(200, "allrooms.gohtml", gin.H{
		"username": UserName,
		"count":    count,
		"wcount":   wishlistcount,
		"rooms":    allrooms,
		"cdate":    cdate,
	})
}

func DateFilter(c *gin.Context) {
	db := database.GetDb()
	session, err := Store.Get(c.Request, "session")
	if err != nil {
		log.Println("Cannot get sessions!")
	}
	useriD := session.Values["userID"]
	userID := fmt.Sprintf("%s", useriD)
	var userinfos models.Users

	db.Raw("SELECT first_name FROM users where email=?", userID).Scan(&userinfos)
	UserName := userinfos.First_Name

	date := c.Param("start")

	a := "available"
	var daterooms []models.Rooms
	db.Where("status=? OR checkoutdate < ?", a, date).Find(&daterooms)

	var UserID int
	db.Raw("SELECT id FROM users WHERE email=?", userID).Scan(&UserID)
	//cart count
	var count int
	db.Raw("SELECT COUNT(user_id) FROM carts WHERE user_id=?", UserID).Scan(&count)
	//wishlist count
	var wishlistcount int
	db.Raw("SELECT COUNT(user_id) FROM wishlists WHERE user_id=?", UserID).Scan(&wishlistcount)

	currentTime := time.Now()
	curdate := currentTime.Format("2006-01-02")
	c.HTML(200, "datefilter.gohtml", gin.H{
		"username": UserName,
		"count":    count,
		"wcount":   wishlistcount,
		"rooms":    daterooms,
		"cdate":date,
		"tdate":curdate,
 	})

}
