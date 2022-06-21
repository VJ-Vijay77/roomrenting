package controllertwo

import (
	"fmt"
	"log"
	"time"

	"github.com/VJ-Vijay77/r4room/pkg/database"
	"github.com/VJ-Vijay77/r4room/pkg/models"
	"github.com/gin-gonic/gin"
)

func AboveFive(c *gin.Context) {
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

	a := "5000"
	var five []models.Rooms
	db.Where("room_price >= ?", a).Find(&five)

	c.HTML(200, "abovefivep.gohtml", gin.H{
		"username": UserName,
		"count":    count,
		"wcount":   wishlistcount,
		"rooms":    five,
		"cdate":    cdate,
	})
}

func AboveSeven(c *gin.Context) {
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

	a := "7000"
	var seven []models.Rooms
	db.Where("room_price >= ?", a).Find(&seven)

	c.HTML(200, "abovesevenp.gohtml", gin.H{
		"username": UserName,
		"count":    count,
		"wcount":   wishlistcount,
		"rooms":    seven,
		"cdate":    cdate,
	})
}

func BelowFive(c *gin.Context) {
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

	a := "5000"
	var bfive []models.Rooms
	db.Where("room_price < ?", a).Find(&bfive)

	c.HTML(200, "belowfivep.gohtml", gin.H{
		"username": UserName,
		"count":    count,
		"wcount":   wishlistcount,
		"rooms":    bfive,
		"cdate":    cdate,
	})
}

//available price filters

func AVAboveFive(c *gin.Context) {
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

	a := "5000"
	s := "available"
	var five []models.Rooms
	db.Where("room_price >= ? AND status=?", a,s).Find(&five)

	c.HTML(200, "avabovefivep.gohtml", gin.H{
		"username": UserName,
		"count":    count,
		"wcount":   wishlistcount,
		"rooms":    five,
		"cdate":    cdate,
	})
}

func AVAboveSeven(c *gin.Context) {
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

	a := "7000"
	s := "available"
	var seven []models.Rooms
	db.Where("room_price >= ? AND status=?", a,s).Find(&seven)

	c.HTML(200, "avabovesevenp.gohtml", gin.H{
		"username": UserName,
		"count":    count,
		"wcount":   wishlistcount,
		"rooms":    seven,
		"cdate":    cdate,
	})
}

func AVBelowFive(c *gin.Context) {
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

	a := "5000"
	s := "available"
	var bfive []models.Rooms
	db.Where("room_price < ? AND status=?", a,s).Find(&bfive)

	c.HTML(200, "avbelowfivep.gohtml", gin.H{
		"username": UserName,
		"count":    count,
		"wcount":   wishlistcount,
		"rooms":    bfive,
		"cdate":    cdate,
	})
}



func BKAboveFive(c *gin.Context) {
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

	a := "5000"
	s := "booked"
	var five []models.Rooms
	db.Where("room_price >= ? AND status=?", a,s).Find(&five)

	c.HTML(200, "bkabovefivep.gohtml", gin.H{
		"username": UserName,
		"count":    count,
		"wcount":   wishlistcount,
		"rooms":    five,
		"cdate":    cdate,
	})
}

func BKAboveSeven(c *gin.Context) {
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

	a := "7000"
	s := "booked"

	var seven []models.Rooms
	db.Where("room_price >= ? AND status=?", a,s).Find(&seven)

	c.HTML(200, "bkabovesevenp.gohtml", gin.H{
		"username": UserName,
		"count":    count,
		"wcount":   wishlistcount,
		"rooms":    seven,
		"cdate":    cdate,
	})
}

func BKBelowFive(c *gin.Context) {
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

	a := "5000"
	s := "booked"

	var bfive []models.Rooms
	db.Where("room_price < ? AND status=?", a,s).Find(&bfive)

	c.HTML(200, "bkbelowfivep.gohtml", gin.H{
		"username": UserName,
		"count":    count,
		"wcount":   wishlistcount,
		"rooms":    bfive,
		"cdate":    cdate,
	})
}

