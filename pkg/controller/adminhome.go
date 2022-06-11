package controller

import (
	"encoding/json"
	"time"

	"github.com/VJ-Vijay77/r4room/pkg/database"
	"github.com/VJ-Vijay77/r4room/pkg/models"
	"github.com/gin-gonic/gin"
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
	db := database.GetDb()
	var categories []models.Category
	db.Raw("SELECT id,category_name FROM category").Scan(&categories)

	c.HTML(200, "addroom.gohtml", gin.H{
		"category": categories,
	})
}

func PAddRoom(c *gin.Context) {
	var ok bool
	//checking session
	// ok := AdminLogedCheck(c)
	// if !ok {
	// 	c.Redirect(303, "/admin_login")
	// 	return
	// }
	db := database.GetDb()
	RoomName := c.PostForm("roomname")
	CoverPicPath := c.PostForm("roompicpath")
	RoomPrice := c.PostForm("roomprice")
	RoomCategory := c.PostForm("roomcategory")
	SubPicPath1 := c.PostForm("roompicpath1")
	SubPicPath2 := c.PostForm("roompicpath2")
	SubPicPath3 := c.PostForm("roompicpath3")
	SubPicPath4 := c.PostForm("roompicpath4")

	var addroom models.Rooms
	db.Raw("SELECT room_name FROM rooms WHERE room_name=?", RoomName).Scan(&addroom)
	if addroom.Room_Name == RoomName {
		ok = false
		k, _ := json.Marshal(ok)
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Write(k)

		return
	}

	addroom.Room_Name = RoomName
	addroom.Room_Price = RoomPrice
	addroom.Category = RoomCategory
	addroom.Cover = CoverPicPath
	addroom.Sub1 = SubPicPath1
	addroom.Sub2 = SubPicPath2
	addroom.Sub3 = SubPicPath3
	addroom.Sub4 = SubPicPath4
	addroom.Status="available"

	db.Select("room_name", "room_price", "category", "cover", "sub1", "sub2", "sub3", "sub4", "sub5","status").Create(&addroom)

	ok = true
	k, _ := json.Marshal(ok)
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.Write(k)
	//c.Redirect(303, "/admin/add_room")
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
	c.HTML(200, "removeroom.gohtml", gin.H{
		"allrooms": room,
	})
}

func DeleteRoom(c *gin.Context) {
	ID := c.Param("ID")

	db := database.GetDb()
	var rooms models.Rooms
	db.Raw("DELETE FROM rooms WHERE id=?", ID).Scan(&rooms)
	time.Sleep(1 * time.Second)
	c.Redirect(303, "/admin/remove_room")
}
