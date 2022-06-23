package controller

import (
	"fmt"

	"github.com/VJ-Vijay77/r4room/pkg/database"
	"github.com/VJ-Vijay77/r4room/pkg/models"
	"github.com/gin-gonic/gin"
)

func UserProfile(c *gin.Context) {
	ok := UserLogedCheck(c)
	if !ok {
		c.Redirect(303,"/user_login")
		return
	}
	db := database.GetDb()
	session, _ := Store.Get(c.Request, "session")
	useriD := session.Values["userID"]
	userID := fmt.Sprintf("%s", useriD)
	var userinfos models.Users
	db.Raw("SELECT first_name,last_name,email,mobile FROM users where email=?", userID).Scan(&userinfos)
	UserName := userinfos.First_Name
	LastName := userinfos.Last_Name
	EmailofUser := userinfos.Email
	MobileOfUser := userinfos.Mobile
	var UserID int
	db.Raw("SELECT id FROM users WHERE email=?", userID).Scan(&UserID)
	//cart count
	var count int
	db.Raw("SELECT COUNT(user_id) FROM carts WHERE user_id=?", UserID).Scan(&count)
	//wishlist count
	var wishlistcount int
	db.Raw("SELECT COUNT(user_id) FROM wishlists WHERE user_id=?", UserID).Scan(&wishlistcount)

	//user address
	var addressofuser models.Useraddress
	var main = "true"
	db.Where("user_id=? AND main=?",UserID,main).Find(&addressofuser)
	
	var wallet models.Wallets
	db.Select("balance").Find(&wallet)

	c.HTML(200, "userprofile.gohtml", gin.H{
		"data": userinfos,
		"username":UserName,
		"count":count,
		"wcount":wishlistcount,
		"lastname":LastName,
		"email":EmailofUser,
		"mobile":MobileOfUser,
		"address":addressofuser,
		"bal":wallet.Balance,
	})
}


func OrderHistory(c *gin.Context) {
	ok := UserLogedCheck(c)
	if !ok {
		c.Redirect(303,"/user_login")
		return
	}

	db := database.GetDb()
	session, _ := Store.Get(c.Request, "session")
	useriD := session.Values["userID"]
	userID := fmt.Sprintf("%s", useriD)
	var userinfos models.Users
	db.Raw("SELECT first_name,last_name,email,mobile FROM users where email=?", userID).Scan(&userinfos)
	UserName := userinfos.First_Name
	// LastName := userinfos.Last_Name
	// EmailofUser := userinfos.Email
	// MobileOfUser := userinfos.Mobile
	var UserID int
	db.Raw("SELECT id FROM users WHERE email=?", userID).Scan(&UserID)
	//cart count
	var count int
	db.Raw("SELECT COUNT(user_id) FROM carts WHERE user_id=?", UserID).Scan(&count)
	//wishlist count
	var wishlistcount int
	db.Raw("SELECT COUNT(user_id) FROM wishlists WHERE user_id=?", UserID).Scan(&wishlistcount)

	var orderinfo []models.Orders
	db.Where("user_id=?",UserID).Find(&orderinfo)

	var roomids []models.Orderedrooms
	db.Where("user_id=?",UserID).Find(&roomids)


	c.HTML(200,"orderhistory.gohtml",gin.H{
		"data": userinfos,
		"username":UserName,
		"count":count,
		"wcount":wishlistcount,
		"orderinfo":orderinfo,
		"roomid":roomids,
		
	})
}

func EditProfile(c *gin.Context) {
	ok := UserLogedCheck(c)
	if !ok {
		c.Redirect(303,"/user_login")
		return
	}

	db := database.GetDb()
	session, _ := Store.Get(c.Request, "session")
	useriD := session.Values["userID"]
	userID := fmt.Sprintf("%s", useriD)
	var userinfos models.Users
	db.Raw("SELECT first_name,last_name,email,mobile FROM users where email=?", userID).Scan(&userinfos)
	UserName := userinfos.First_Name

	var UserID int
	db.Raw("SELECT id FROM users WHERE email=?", userID).Scan(&UserID)
	//cart count
	var count int
	db.Raw("SELECT COUNT(user_id) FROM carts WHERE user_id=?", UserID).Scan(&count)
	//wishlist count
	var wishlistcount int
	db.Raw("SELECT COUNT(user_id) FROM wishlists WHERE user_id=?", UserID).Scan(&wishlistcount)

	



	c.HTML(200,"editprofile.gohtml",gin.H{
		"data": userinfos,
		"username":UserName,
		"count":count,
		"wcount":wishlistcount,
		"ID":UserID,
	})
}


func EditProfileTwo(c *gin.Context) {
	ok := UserLogedCheck(c)
	if !ok {
		c.Redirect(303,"/user_login")
		return
	}

	db := database.GetDb()
	session, _ := Store.Get(c.Request, "session")
	useriD := session.Values["userID"]
	userID := fmt.Sprintf("%s", useriD)
	var userinfos models.Users
	db.Raw("SELECT first_name,last_name,email,mobile FROM users where email=?", userID).Scan(&userinfos)
	UserName := userinfos.First_Name

	var UserID int
	db.Raw("SELECT id FROM users WHERE email=?", userID).Scan(&UserID)
	//cart count
	var count int
	db.Raw("SELECT COUNT(user_id) FROM carts WHERE user_id=?", UserID).Scan(&count)
	//wishlist count
	var wishlistcount int
	db.Raw("SELECT COUNT(user_id) FROM wishlists WHERE user_id=?", UserID).Scan(&wishlistcount)

	



	c.HTML(200,"editprofiletwo.gohtml",gin.H{
		"data": userinfos,
		"username":UserName,
		"count":count,
		"wcount":wishlistcount,
		"ID":UserID,
	})
}