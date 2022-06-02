package controller

import (
	"github.com/VJ-Vijay77/r4room/pkg/database"
	"github.com/VJ-Vijay77/r4room/pkg/models"
	"github.com/gin-gonic/gin"
	"tawesoft.co.uk/go/dialog"
)

func UpdateFUser(c *gin.Context) {
	FName := c.Request.FormValue("updateuserdata")
	if FName==""{
		dialog.Alert("Enter data!!")
		c.Redirect(303,"/admin/user_management")
		return
	}
	var data models.Users
	ID := c.Param("ID")
	db := database.GetDb()
	db.Raw("UPDATE users SET first_name=? WHERE id=?", FName, ID).Scan(&data)
	c.Redirect(303, "/admin/user_management")

}
func UpdateLUser(c *gin.Context) {

	LName := c.Request.FormValue("updateuserdata")
	if LName==""{
		dialog.Alert("Enter data!!")
		c.Redirect(303,"/admin/user_management")
		return
	}
	
	var data models.Users
	ID := c.Param("ID")
	db := database.GetDb()
	db.Raw("UPDATE users SET last_name=? WHERE id=?", LName, ID).Scan(&data)
	c.Redirect(303, "/admin/user_management")

}
