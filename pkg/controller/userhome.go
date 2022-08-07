package controller

import (
	"fmt"
	"time"

	"github.com/VJ-Vijay77/r4room/pkg/database"
	"github.com/VJ-Vijay77/r4room/pkg/models"
	"github.com/gin-gonic/gin"
)

//User home page
func UserHome(c *gin.Context) {
		
	//checking session
	ok := UserLogedCheck(c)
	if !ok {
		c.HTML(200, "userhome.gohtml", nil)
		return
	}
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

	var rooms []models.Rooms
	db.Find(&rooms)

	time := time.Now().Format("2006-01-02")
	var roomsupdation models.Rooms
	var orderedupdation models.Orderedrooms
	var availablenow models.Rooms
	var a = "checkedin"
	for _,i := range rooms{
		if i.Checkoutdate < time{
	db.Raw("UPDATE rooms SET status='available' WHERE id=?", i.ID).Scan(&roomsupdation)
	db.Raw("UPDATE orderedrooms SET status='checkedout' WHERE roomid=? AND status=?", i.ID,a).Scan(&orderedupdation)
	db.Raw("UPDATE rooms SET checkoutdate='availablenow' WHERE id=?", i.ID).Scan(&availablenow)
    			
		}
	}

	var allrooms []models.Rooms
	var allroomses []models.Rooms
	var adroom models.Rooms
	db.Find(&adroom).Where("id=?",65)
	db.Limit(4).Find(&allrooms)
	db.Limit(4).Offset(4).Find(&allroomses)
	
	c.HTML(200, "userhome.gohtml", gin.H{
		"data":     userinfos,
		"username": UserName,
		"count":    count,
		"wcount":   wishlistcount,
		"ad":adroom,
		"rooms":allrooms,
		"offrooms":allroomses,
	})
}
