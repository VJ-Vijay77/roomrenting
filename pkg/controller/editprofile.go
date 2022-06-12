package controller

import (
	"fmt"
	"strconv"

	"github.com/VJ-Vijay77/r4room/pkg/database"
	"github.com/VJ-Vijay77/r4room/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
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
	db := database.GetDb()

	lastname := c.PostForm("lastnamef")
	userID := c.Param("ID")
	var lastnameupdation models.Users
	db.Raw("UPDATE users SET last_name=? WHERE id=?", lastname, userID).Scan(&lastnameupdation)
	dialog.Alert("Last name updated Successfully.")
	c.Redirect(303, "/user/edit_profile")

}

func EditProfileMobile(c *gin.Context) {
	db := database.GetDb()

	mobile := c.PostForm("mobilef")
	userID := c.Param("ID")
	var mobileupdation models.Users
	db.Raw("UPDATE users SET mobile=? WHERE id=?", mobile, userID).Scan(&mobileupdation)
	dialog.Alert("Mobile number updated Successfully.")
	c.Redirect(303, "/user/edit_profile")
}

func EditProfileEmail(c *gin.Context) {
	db := database.GetDb()

	session, _ := Store.Get(c.Request, "session")
	oldemail := fmt.Sprintf("%v", session.Values["userID"])

	email := c.PostForm("emailf")
	userID := c.Param("ID")
	if oldemail == email {
		dialog.Alert("You entered same email! Try with different one!")
		c.Redirect(303, "/user/edit_profile")
		return
	}

	var emailupdation models.Users
	db.Raw("UPDATE users SET email=? WHERE id=?", email, userID).Scan(&emailupdation)

	dialog.Alert("Email updated Successfully.\n\t\tNOTE!!\nYour login email will be now %s\nYou may need to login again now!", email)
	session.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	}
	session.Save(c.Request, c.Writer)
	c.Redirect(303, "/user_login")

}

func EditProfileAddress(c *gin.Context) {
	uID := c.Param("ID")
	Id,_ := strconv.Atoi(uID)
	 db := database.GetDb()
	HouseName := c.PostForm("housenamef")
	Place := c.PostForm("placef")
	Mobile := c.PostForm("mobile2f")
	State := c.PostForm("statef")
	uPIN := c.PostForm("pinf")
	PIN,_ := strconv.Atoi(uPIN)


	
	db.AutoMigrate(&models.Useraddress{})
	var userid models.Useraddress
	db.Raw("SELECT user_id FROM useraddresses WHERE user_id=?",Id).Scan(&userid)
	if userid.User_Id == Id {
		db.Raw("UPDATE useraddresses SET housename=? WHERE user_id=?",HouseName,Id).Scan(&userid)
		db.Raw("UPDATE useraddresses SET place=? WHERE user_id=?",Place,Id).Scan(&userid)
		db.Raw("UPDATE useraddresses SET mobile=? WHERE user_id=?",Mobile,Id).Scan(&userid)
		db.Raw("UPDATE useraddresses SET state=? WHERE user_id=?",State,Id).Scan(&userid)
		db.Raw("UPDATE useraddresses SET pin=? WHERE user_id=?",PIN,Id).Scan(&userid)


	}else{
		
		userid.User_Id=Id
		userid.Housename=HouseName
		userid.Place=Place
		userid.Mobile=Place
		userid.State=State
		userid.PIN=PIN
		db.Select("user_id","housename","place","mobile","state","pin").Create(&userid)
	}


}


func AddressTwo (c *gin.Context) {
	
}
