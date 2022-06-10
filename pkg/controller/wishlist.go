package controller

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/VJ-Vijay77/r4room/pkg/database"
	"github.com/VJ-Vijay77/r4room/pkg/models"
	"github.com/gin-gonic/gin"
)

func Wishlish(c *gin.Context) {

	session, _ := Store.Get(c.Request, "session")
	email := session.Values["userID"]
	user_ID := fmt.Sprintf("%v", email)

	db := database.GetDb()
	var UserID int
	db.Raw("SELECT id FROM users WHERE email=?", user_ID).Scan(&UserID)

	// counter setting
	var count int
	db.Raw("SELECT COUNT(user_id) FROM carts WHERE user_id=?", UserID).Scan(&count)

	var userinfos models.Users
	db.Raw("SELECT first_name FROM users where email=?", user_ID).Scan(&userinfos)
	UserName := userinfos.First_Name
	var wishlistitems []models.Wishlist_Infos
	db.Raw("SELECT wishlists.wishlistid,users.id,rooms.room_name,rooms.cover,rooms.room_price,rooms.category FROM wishlists INNER JOIN rooms ON rooms.id=wishlists.wishroomsid INNER JOIN users ON wishlists.user_id=users.id WHERE user_id=?", UserID).Scan(&wishlistitems)
	c.HTML(200, "wishlist.gohtml", gin.H{
		"wishlist":     wishlistitems,
		"username": UserName,
		"count":    count,
	})

}

func AddToWishlist(c *gin.Context) {

	ok := UserLogedCheck(c)
	if !ok {
		var not = "login"
		k, _ := json.Marshal(not)
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Write(k)
		return

	}


	db := database.GetDb()
	session, _ := Store.Get(c.Request, "session")
	userID := session.Values["userID"]
	userEmail := fmt.Sprintf("%v", userID)


	I := c.Param("RID")
	RoomID, _ := strconv.Atoi(I)


	var usersID int
	//  var roominfo models.Rooms
	
	db.Raw("SELECT id FROM users WHERE email=?", userEmail).Scan(&usersID)
	
	var wishlist models.Wishlists
	db.Where("wishroomsid=? AND user_id=?", RoomID, usersID).Find(&wishlist)
	
	if wishlist.Wishroomsid == RoomID && wishlist.User_ID == usersID {
		var room = "sameroom"
		k, _ := json.Marshal(room)
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Write(k)
		return
	}
	var wish models.Wishlists

	wish.Wishroomsid = RoomID
	wish.User_ID = usersID
	
	db.Select("wishroomsid", "user_id").Create(&wish)
	var note = "added"
	k, _ := json.Marshal(note)
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.Write(k)

}

func DeleteFromWishlist(c *gin.Context) {
	CID := c.Param("CID")
	CartId, _ := strconv.Atoi(CID)
	db := database.GetDb()
	var cart models.Carts
	db.Raw("DELETE FROM carts WHERE cartsid=?", CartId).Scan(&cart)
	time.Sleep(1 * time.Second)
	c.Redirect(303, "/user/cart")

}
