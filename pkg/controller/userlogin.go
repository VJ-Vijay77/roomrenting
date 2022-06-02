package controller

import (
	"log"
	"os"

	"github.com/VJ-Vijay77/r4room/pkg/database"

	"github.com/VJ-Vijay77/r4room/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"tawesoft.co.uk/go/dialog"
)

//setting sessions using gorilla sesions
var Store = sessions.NewCookieStore([]byte(os.Getenv("session")))

//User login get and post methods
func UserLogin(c *gin.Context) {

	//checking session
	ok := UserLogedCheck(c)
	if ok {
		c.Redirect(303, "/user_home")
		return
	}

	c.HTML(200, "userlogin.gohtml", nil)
}

//validating user credentials from database
func PUserLogin(c *gin.Context) {
	db := database.GetDb()
	Email := c.PostForm("useremail")
	Password := c.PostForm("userpassword")

	var ds models.Users
	db.Raw("SELECT email,password,block_status FROM users WHERE email=? AND password=?", Email, Password).Scan(&ds)
	if Email != ds.Email || Password != ds.Password {
		dialog.Alert("Username or Password is Incorrect\n\t\tTry again!")
		c.Redirect(303, "/user_login")
		return
	}
	if !ds.Block_Status {
		dialog.Alert("You have been blocked by Admin!\n  contact:vijayadmin@admin")
		c.Redirect(303, "/user_login")
		return
	}

	//creating sessions for the user
	session, err := Store.Get(c.Request, "session")
	if err != nil {
		log.Println("Error getting session !!")
	}
	session.Values["userID"] = Email
	session.Save(c.Request, c.Writer)
	c.Redirect(303, "/user_home")

}

func UserLogout(c *gin.Context) {
	session, err := Store.Get(c.Request, "session")
	if err != nil {
		log.Println("couldnt fetch sessions!")
	}
	session.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	}
	sessions.Save(c.Request, c.Writer)
	c.Redirect(303, "/user_login")
}
