package controller

import (
	"github.com/VJ-Vijay77/r4room/pkg/database"
	"github.com/VJ-Vijay77/r4room/pkg/models"
	"github.com/gin-gonic/gin"
	"tawesoft.co.uk/go/dialog"
)

func EditProfileFirstName(c *gin.Context) {
	db := database.GetDb()

	firstname := c.PostForm("firstnamef")
	userID := c.Param("ID")
	var firstnameupdation models.Users
	db.Raw("UPDATE users SET first_name=? WHERE id=?", firstname, userID).Scan(&firstnameupdation)
	dialog.Alert("First name updated Successfully.")
	c.Redirect(303, "/user/edit_profile")
}

func EditProfileLastName(c *gin.Context) {

}

func EditProfileMobile(c *gin.Context) {

}

func EditProfileEmail(c *gin.Context) {

}

func EditProfileAddress(c *gin.Context) {

}
