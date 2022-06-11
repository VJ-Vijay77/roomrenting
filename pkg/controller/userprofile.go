package controller

import (
	"fmt"

	"github.com/VJ-Vijay77/r4room/pkg/database"
	"github.com/VJ-Vijay77/r4room/pkg/models"
	"github.com/gin-gonic/gin"
)

func UserProfile(c *gin.Context) {
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

	c.HTML(200, "userprofile.gohtml", gin.H{
		"data": userinfos,
		"username":UserName,
		"count":count,
		"wcount":wishlistcount,
		"lastname":LastName,
		"email":EmailofUser,
		"mobile":MobileOfUser,
	})
}


func OrderHistory(c *gin.Context) {

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



	c.HTML(200,"orderhistory.gohtml",gin.H{
		"data": userinfos,
		"username":UserName,
		"count":count,
		"wcount":wishlistcount,
		
	})
}