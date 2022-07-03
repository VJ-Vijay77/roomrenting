package controllertwo

import (
	"encoding/json"
	"fmt"
	"strconv"
	
	"github.com/VJ-Vijay77/r4room/pkg/controller"
	"github.com/VJ-Vijay77/r4room/pkg/database"
	"github.com/VJ-Vijay77/r4room/pkg/models"
	"github.com/gin-gonic/gin"
)

var Store = controller.Store

func Settings(c *gin.Context) {
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

	c.HTML(200, "settings.gohtml", gin.H{
		"data":     userinfos,
		"username": UserName,
		"count":    count,
		"wcount":   wishlistcount,
	})
}

func ChangePassword(c *gin.Context) {
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

	c.HTML(200, "changepassword.gohtml", gin.H{
		"data":     userinfos,
		"username": UserName,
		"ID":       UserID,
		"count":    count,
		"wcount":   wishlistcount,
	})
}

func PChangePassword(c *gin.Context) {
	db := database.GetDb()
	var ok string
	uID := c.Param("ID")
	ID, _ := strconv.Atoi(uID)
	oldpassword := c.PostForm("oldpassword")
	newpassword := c.PostForm("newpassword")
	renewpassword := c.PostForm("renewpassword")

	// fmt.Println(uID,db.Name(),oldpassword,newpassword,renewpassword)

	var users models.Users
	db.Raw("SELECT password FROM users WHERE id=?", ID).Scan(&users)

	if oldpassword != users.Password {
		ok = "wrongpassword"
		k, _ := json.Marshal(ok)
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Write(k)
		return
	}
	if oldpassword == newpassword {
		ok = "same"
		k, _ := json.Marshal(ok)
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Write(k)
		return
	}

	if newpassword != renewpassword {
		ok = "different"
		k, _ := json.Marshal(ok)
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Write(k)
		return
	}

	var passwordupdate models.Users
	db.Raw("UPDATE users SET password=? WHERE id=?", newpassword, ID).Scan(&passwordupdate)

	ok = "done"
	k, _ := json.Marshal(ok)
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.Write(k)

}
