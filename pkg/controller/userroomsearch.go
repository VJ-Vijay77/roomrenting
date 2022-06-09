package controller

import (
	"fmt"
	"log"

	"github.com/VJ-Vijay77/r4room/pkg/database"
	"github.com/VJ-Vijay77/r4room/pkg/models"
	"github.com/gin-gonic/gin"
)

func SearchRooms(c *gin.Context) {
	db := database.GetDb()

	session, err := Store.Get(c.Request, "session")
	if err !=nil {log.Println("Cannot get sessions!")}
	useriD := session.Values["userID"]
	userID := fmt.Sprintf("%s", useriD)
	var userinfos models.Users

	db.Raw("SELECT first_name FROM users where email=?", userID).Scan(&userinfos)

	UserName := userinfos.First_Name
	var rooms []models.Rooms
	
	db.Order("room_name").Find(&rooms)
	var UserID int
	db.Raw("SELECT id FROM users WHERE email=?", userID).Scan(&UserID)
	var count int
	db.Raw("SELECT COUNT(user_id) FROM carts WHERE user_id=?",UserID).Scan(&count)
	
	
	c.HTML(200, "roomsearchpage.gohtml", gin.H{
		"rooms":    rooms,
		"username": UserName,
		"count":count,
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
	db.Where("id=?",ID).Find(&roominfo)

	
	c.HTML(200, "roominfo.gohtml", gin.H{
		 "roomname":     roominfo.Room_Name,
		 "roomid":       roominfo.ID,
		 "roomcategory": roominfo.Category,
		 "roomprice":    roominfo.Room_Price,
		 "cover":roominfo.Cover,
		 "sub1":roominfo.Sub1,
		 "sub2":roominfo.Sub2,
		 "sub3":roominfo.Sub3,
		 "sub4":roominfo.Sub4,
		 "sub5":roominfo.Sub5,
		"username":UserName,
		//"info":roominfo,
		
	})
}
