package controllertwo

import (
	"fmt"
	"log"
	"time"

	"github.com/VJ-Vijay77/r4room/pkg/database"
	"github.com/VJ-Vijay77/r4room/pkg/models"
	"github.com/gin-gonic/gin"
)

func SingleRoomFilter(c *gin.Context) {
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

	a := "single"
	var single []models.Rooms
	db.Where("category=?", a).Find(&single)

	c.HTML(200, "filtersingle.gohtml", gin.H{
		"username": UserName,
		"count":    count,
		"wcount":   wishlistcount,
		"rooms":    single,
		"cdate":    cdate,
	})
}

func DoubleRoomFilter(c *gin.Context) {
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

	a := "double"
	var double []models.Rooms
	db.Where("category=?", a).Find(&double)

	c.HTML(200, "doublefilter.gohtml", gin.H{
		"username": UserName,
		"count":    count,
		"wcount":   wishlistcount,
		"rooms":    double,
		"cdate":    cdate,
	})
}



func AVSingleRoomFilter(c *gin.Context) {
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

	a := "single"
	av := "available"
	var single []models.Rooms
	db.Where("category=? AND status=?", a,av).Find(&single)

	c.HTML(200, "avsingle.gohtml", gin.H{
		"username": UserName,
		"count":    count,
		"wcount":   wishlistcount,
		"rooms":    single,
		"cdate":    cdate,
	})
}

func AVDoubleRoomFilter(c *gin.Context) {
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

	a := "double"
	av := "available"
	var double []models.Rooms
	db.Where("category=? AND status=?", a,av).Find(&double)

	c.HTML(200, "avdouble.gohtml", gin.H{
		"username": UserName,
		"count":    count,
		"wcount":   wishlistcount,
		"rooms":    double,
		"cdate":    cdate,
	})
}




func BKSingleRoomFilter(c *gin.Context) {
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

	a := "single"
	av := "booked"
	var single []models.Rooms
	db.Where("category=? AND status=?", a,av).Find(&single)

	c.HTML(200, "bksingle.gohtml", gin.H{
		"username": UserName,
		"count":    count,
		"wcount":   wishlistcount,
		"rooms":    single,
		"cdate":    cdate,
	})
}

func BKDoubleRoomFilter(c *gin.Context) {
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

	a := "double"
	av := "booked"
	var double []models.Rooms
	db.Where("category=? AND status=?", a,av).Find(&double)

	c.HTML(200, "bkdouble.gohtml", gin.H{
		"username": UserName,
		"count":    count,
		"wcount":   wishlistcount,
		"rooms":    double,
		"cdate":    cdate,
	})
}
