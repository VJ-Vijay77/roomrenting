package controller

import (
	"github.com/VJ-Vijay77/r4room/pkg/database"
	"github.com/VJ-Vijay77/r4room/pkg/models"
	"github.com/gin-gonic/gin"
	"tawesoft.co.uk/go/dialog"
)

//User login get and post methods
func UserLogin(c *gin.Context) {
	c.HTML(200, "userlogin.gohtml", nil)
}

//validating user credentials from database
func PUserLogin(c *gin.Context) {
	db := database.GetDb()
	Email := c.PostForm("useremail")
	Password := c.PostForm("userpassword")
	
	var ds models.Users
	db.Raw("SELECT email,password FROM users WHERE email=? AND password=?", Email,Password).Scan(&ds)
	if Email != ds.Email || Password != ds.Password {
		dialog.Alert("Username or Password is Incorrect\n\t\tTry again!")
		c.Redirect(303,"/user_login")
		return
	}
	c.Redirect(303, "/user_home")

}


