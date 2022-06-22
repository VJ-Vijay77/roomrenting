package controller

import (
	"encoding/json"
	"path/filepath"
	"strconv"
	"time"

	"github.com/VJ-Vijay77/r4room/pkg/database"
	"github.com/VJ-Vijay77/r4room/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	// CoverPicPath := c.PostForm("roompicpath")
	CoverPicPath, _ := c.FormFile("roompicpath")
	extension := filepath.Ext(CoverPicPath.Filename)
	Cover := uuid.New().String() + extension
	c.SaveUploadedFile(CoverPicPath, "./public/"+Cover)

	RoomPrice := c.PostForm("roomprice")
	RoomCost ,_ := strconv.Atoi(RoomPrice)
	RoomCategory := c.PostForm("roomcategory")
	RoomDescription := c.PostForm("description")
	SubPicPath1, _ := c.FormFile("roompicpath1")
	extension = filepath.Ext(SubPicPath1.Filename)
	SubPic1 := uuid.New().String() + extension
	c.SaveUploadedFile(SubPicPath1, "./public/"+SubPic1)
	// SubPicPath1 := c.PostForm("roompicpath1")
	SubPicPath2, _ := c.FormFile("roompicpath2")
	extension = filepath.Ext(SubPicPath2.Filename)
	SubPic2 := uuid.New().String() + extension
	c.SaveUploadedFile(SubPicPath2, "./public/"+SubPic2)

	// SubPicPath2 := c.PostForm("roompicpath2")
	SubPicPath3, _ := c.FormFile("roompicpath3")
	extension = filepath.Ext(SubPicPath3.Filename)
	SubPic3 := uuid.New().String() + extension
	c.SaveUploadedFile(SubPicPath3, "./public/"+SubPic3)

	// SubPicPath3 := c.PostForm("roompicpath3")
	SubPicPath4, _ := c.FormFile("roompicpath4")
	extension = filepath.Ext(SubPicPath4.Filename)
	SubPic4 := uuid.New().String() + extension
	c.SaveUploadedFile(SubPicPath4, "./public/"+SubPic4)

	// SubPicPath4 := c.PostForm("roompicpath4")

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
	addroom.Room_Price = RoomCost
	addroom.Category = RoomCategory
	addroom.Cover = Cover
	addroom.Sub1 = SubPic1
	addroom.Sub2 = SubPic2
	addroom.Sub3 = SubPic3
	addroom.Sub4 = SubPic4
	addroom.Status = "available"
	addroom.Description = RoomDescription

	db.Select("room_name", "room_price", "category", "cover", "sub1", "sub2", "sub3", "sub4", "sub5", "status","description").Create(&addroom)

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
