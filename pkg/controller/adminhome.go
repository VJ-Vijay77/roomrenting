package controller

import (
	"fmt"

	"github.com/VJ-Vijay77/r4room/pkg/database"
	"github.com/VJ-Vijay77/r4room/pkg/models"
	"github.com/gin-gonic/gin"
	"tawesoft.co.uk/go/dialog"
)

//rendering admin homepage
func AdminHome(c *gin.Context) {

	//checking session
	ok := AdminLogedCheck(c)
	if !ok {
		c.Redirect(303, "/admin_login")
		return
	}

	c.HTML(200, "adminhome.gohtml", nil)
}

func UserManagement(c *gin.Context) {

	//checking session
	ok := AdminLogedCheck(c)
	if !ok {
		c.Redirect(303, "/admin_login")
		return
	}

	c.HTML(200, "usermanagement.gohtml", nil)
}

//rendering all users page
func AllUsers(c *gin.Context) {

	//checking session
	ok := AdminLogedCheck(c)
	if !ok {
		c.Redirect(303, "/admin_login")
		return
	}

	db := database.GetDb()
	var data []models.Users

	//ordering all data by id
	db.Order("id").Find(&data)

	c.HTML(200, "allusers.gohtml", gin.H{
		"data": data,
	})
}

//rendering add room page
func AddRoom(c *gin.Context) {

	//checking session
	ok := AdminLogedCheck(c)
	if !ok {
		c.Redirect(303, "/admin_login")
		return
	}

	c.HTML(200, "addroom.gohtml", nil)
}

func PAddRoom(c *gin.Context) {

	//checking session
	// ok := AdminLogedCheck(c)
	// if !ok {
	// 	c.Redirect(303, "/admin_login")
	// 	return
	// }
	db := database.GetDb()
	RoomName := c.PostForm("roomname")
	RoomPicPath := c.PostForm("roompicpath")
	RoomPrice := c.PostForm("roomprice")
	RoomCategory := c.PostForm("roomcategory")

	var addroom models.Rooms

	addroom.Room_Name = RoomName
	addroom.Room_Photo_Path = RoomPicPath
	addroom.Room_Price = RoomPrice
	addroom.Category = RoomCategory

	db.Select("room_name","room_photo_path","room_price", "category").Create(&addroom)

	//    if  {
	// 			dialog.Alert("Something Went Wrong!")
	// 			c.Redirect(303,"/admin/add_room")
	// 		}
	dialog.Alert("Room Added Successfully!")
	c.Redirect(303, "/admin/add_room")
}

//rendering remove room page
func RemoveRoom(c *gin.Context) {

	//checking session
	ok := AdminLogedCheck(c)
	if !ok {
		c.Redirect(303, "/admin_login")
		return
	}
	db := database.GetDb()
	var room []models.Rooms
	db.Order("id").Find(&room)
	c.HTML(200,"removeroom.gohtml",gin.H{
		"allrooms":room,
	})
}

func DeleteRoom(c *gin.Context){
	ID := c.Param("ID")
	fmt.Println(ID)
	db := database.GetDb()
	var rooms models.Rooms
	db.Raw("DELETE FROM rooms WHERE id=?",ID).Scan(&rooms)
	c.Redirect(303,"/admin/remove_room")
}