package controller

import (
	"fmt"

	"github.com/VJ-Vijay77/r4room/pkg/database"
	"github.com/VJ-Vijay77/r4room/pkg/models"
	"github.com/gin-gonic/gin"
)

func SearchRooms(c *gin.Context) {
	db := database.GetDb()

	session, _ := Store.Get(c.Request, "session")
	useriD := session.Values["userID"]
	userID := fmt.Sprintf("%s", useriD)
	var userinfos models.Users

	db.Raw("SELECT first_name FROM users where email=?", userID).Scan(&userinfos)

	UserName := userinfos.First_Name
	var rooms []models.Rooms
	db.Order("room_name").Find(&rooms)
	c.HTML(200, "roomsearchpage.gohtml", gin.H{
		"rooms":    rooms,
		"username": UserName,
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
	db.Raw("SELECT room_name,id,room_price,category FROM rooms WHERE id=?", ID).Scan(&roominfo)
	c.HTML(200, "roominfo.gohtml", gin.H{
		"roomname":     roominfo.Room_Name,
		"roomid":       roominfo.ID,
		"roomcategory": roominfo.Category,
		"roomprice":    roominfo.Room_Price,
		"username":UserName,
	})
}
