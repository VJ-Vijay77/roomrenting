package controllertwo

import (
	"fmt"
	"strconv"
	"time"

	"github.com/VJ-Vijay77/r4room/pkg/database"
	"github.com/VJ-Vijay77/r4room/pkg/models"
	"github.com/gin-gonic/gin"
)

func CheckingsList(c *gin.Context) {
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

	var bookings []models.BookingDetails
	
db.Raw("SELECT user_id,rooms.id,rooms.room_name,orderedrooms.status,rooms.cover,rooms.category,rooms.room_price FROM orderedrooms INNER JOIN rooms ON rooms.id=orderedrooms.roomid WHERE user_id=?",UserID).Scan(&bookings)
	
	
	
	c.HTML(200, "checkingslist.gohtml", gin.H{
		"data":     userinfos,
		"username": UserName,
		"count":    count,
		"wcount":   wishlistcount,
		"UID":UserID,
		"bookings":bookings,
	})
}



func CheckingOut(c *gin.Context){
	RoomId := c.Param("RID")
	RID,_ := strconv.Atoi(RoomId)
	userId := c.Param("UID")
	UID,_ := strconv.Atoi(userId)
	db := database.GetDb()

	var roomupdations models.Orderedrooms
	db.Raw("UPDATE orderedrooms SET status='checkedout' WHERE roomid=? AND user_id=?",RID,UID).Scan(&roomupdations)

	var roomstatus models.Rooms
	db.Raw("UPDATE rooms SET status='available' WHERE id=?",RID).Scan(&roomstatus)

	time.Sleep(2 *time.Second)

	c.Redirect(303,"/user/checkings")
}