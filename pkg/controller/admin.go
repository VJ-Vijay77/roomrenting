package controller

import (
	"log"

	"github.com/VJ-Vijay77/r4room/pkg/database"
	"github.com/VJ-Vijay77/r4room/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"tawesoft.co.uk/go/dialog"
)

//admin login page
func AdminLogin(c *gin.Context) {

		//checking session
		ok := AdminLogedCheck(c)
		if ok {
			c.Redirect(303, "/admin_home")
			return
		}	

	c.HTML(200, "adminlogin.gohtml", gin.H{})
}

//validating admin credentials from database
func PAdminLogin(c *gin.Context) {
	db := database.GetDb()
	Email := c.PostForm("email")
	Password := c.PostForm("password")

	var status bool
	var admin models.Users
	db.Raw("SELECT email,password,role FROM users WHERE email=?", Email).Scan(&admin)

	if admin.Email == Email && admin.Password == Password && admin.Role == "admin" {
		status = true
	}
	if !status {
		dialog.Alert("Wrong username or password !!\n\t\tTry again.")
		c.Redirect(303, "/admin_login")
		return
	}
	
	//setting session for admin
	session, err := Store.Get(c.Request, "adminsession")
	if err != nil {
		log.Println("Error getting session !!")
	}
	session.Values["userID"] = Email
	session.Save(c.Request, c.Writer)
	c.Redirect(303, "/admin_home")

}

//admin logout
func AdminLogout(c *gin.Context) {
	session, err := Store.Get(c.Request, "adminsession")
	if err != nil {
		log.Println("couldnt fetch sessions!")
	}
	session.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	}
	sessions.Save(c.Request,c.Writer)
	c.Redirect(303,"/admin_login")
}

