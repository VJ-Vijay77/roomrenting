package controller

import (
	"github.com/VJ-Vijay77/r4room/pkg/database"
	"github.com/VJ-Vijay77/r4room/pkg/models"
	"github.com/gin-gonic/gin"
)

//rendering admin homepage
func AdminHome(c *gin.Context) {
	c.HTML(200, "adminhome.gohtml", nil)
}

func UserManagement(c *gin.Context){
	c.HTML(200,"usermanagement.gohtml",nil)
}



//rendering all users page
func AllUsers(c *gin.Context) {
	db := database.GetDb()
	var data []models.Users
	
	//ordering all data by id
	db.Order("id").Find(&data)

	c.HTML(200, "allusers.gohtml", gin.H{
		"data":data,
	})
}

//rendering add room page
func AddRoom(c *gin.Context) {
	c.HTML(200, "addroom.gohtml", nil)
}

//rendering remove room page
func RemoveRoom(c *gin.Context) {
	c.HTML(200, "removeroom.gohtml", nil)
}
