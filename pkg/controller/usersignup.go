package controller

import (
	"github.com/VJ-Vijay77/r4room/pkg/database"
	"github.com/VJ-Vijay77/r4room/pkg/models"
	"github.com/gin-gonic/gin"
	"tawesoft.co.uk/go/dialog"
)


//user signup get and post methods
func UserSignup(c *gin.Context) {
	c.HTML(200, "usersignup.gohtml", nil)
}

//validating user credentials from database
func PUserSignup(c *gin.Context) {
	db := database.GetDb()
	FirstName := c.PostForm("firstname")
	LastName := c.PostForm("lastname")
	Email := c.PostForm("email")
	Password := c.PostForm("password")
	Mobile := c.PostForm("phone")

	var users models.Users

	db.Raw("SELECT email FROM users WHERE email=?",Email).Scan(&users)
	if users.Email == Email{
		dialog.Alert("Hey %s ,The Email is already taken!!",FirstName)
		c.Redirect(303,"/user_signup")
		return
	}
	db.Raw("SELECT mobile FROM users WHERE mobile=?",Mobile).Scan(&users)
	if users.Mobile == Mobile{
		dialog.Alert("Hey %s ,The Mobile is already taken!!",FirstName)
		c.Redirect(303,"/user_signup")
		return
	}


	
		users.First_Name=FirstName
		users.Last_Name=LastName
		users.Email=Email
		users.Password=Password
		users.Block_Status=true
		users.Mobile=Mobile
		
    db.Select("first_name","last_name","email","password","block_status","mobile").Create(&users)
	dialog.Alert("Account Create Successfully!,\n\t\tClick Ok to Login.")
	
	c.Redirect(303,"/user_login")
	
}