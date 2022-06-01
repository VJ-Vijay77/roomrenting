package controller

import (
	"github.com/VJ-Vijay77/r4room/pkg/database"
	"github.com/VJ-Vijay77/r4room/pkg/models"
	"github.com/gin-gonic/gin"
	"tawesoft.co.uk/go/dialog"
)

//admin login page
func AdminLogin(c *gin.Context) {
	c.HTML(200, "adminlogin.gohtml", gin.H{})
}

//validating admin credentials from database
func PAdminLogin(c *gin.Context) {
	db := database.GetDb()
	Email := c.PostForm("email")
	Password := c.PostForm("password")
	
	var status bool
	var admin models.Users
	db.Raw("SELECT email,password,role FROM users WHERE email=?",Email).Scan(&admin)

	if (admin.Email==Email && admin.Password==Password && admin.Role=="admin"){
		status = true
	}
	if !status {
		dialog.Alert("Wrong username or password !!\n\t\tTry again.")
		c.Redirect(303,"/admin_login")
		return
	}
	c.Redirect(303,"/admin_home")

}

