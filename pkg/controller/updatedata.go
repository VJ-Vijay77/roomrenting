package controller

import (
	"github.com/VJ-Vijay77/r4room/pkg/database"
	"github.com/VJ-Vijay77/r4room/pkg/models"
	"github.com/gin-gonic/gin"
	"tawesoft.co.uk/go/dialog"
)
//updating first name of user
func UpdateFUser(c *gin.Context) {
	
	FName := c.Request.FormValue("updateuserdata")
	if FName==""{
		dialog.Alert("Enter data!!")
		c.Redirect(303,"/admin/user_management/allusers")
		return
	}
	var data models.Users
	ID := c.Param("ID")
	if ID == "1"{
		dialog.Alert("Cant do anything,User has Admin privilages!")
		c.Redirect(303,"/admin/user_management/allusers")
		return
	}
	db := database.GetDb()
	db.Raw("UPDATE users SET first_name=? WHERE id=?", FName, ID).Scan(&data)
	c.Redirect(303, "/admin/user_management/allusers")

}

//updating last name of user
func UpdateLUser(c *gin.Context) {

	LName := c.Request.FormValue("updateuserdata")
	if LName==""{
		dialog.Alert("Enter data!!")
		c.Redirect(303,"/admin/user_management/allusers")
		return
	}

	var data models.Users
	ID := c.Param("ID")
	if ID == "1"{
		dialog.Alert("Cant do anything,User has Admin privilages!")
		c.Redirect(303,"/admin/user_management/allusers")
		return
	}
	db := database.GetDb()
	db.Raw("UPDATE users SET last_name=? WHERE id=?", LName, ID).Scan(&data)
	c.Redirect(303, "/admin/user_management/allusers")

}

//setting block status of user to false
func BlockUser(c *gin.Context) {
	var user models.Users
	ID := c.Param("ID")
	if ID == "1"{
		dialog.Alert("Cant do anything,User has Admin privilages!")
		c.Redirect(303,"/admin/user_management/allusers")
		return
	}
	db := database.GetDb()
	db.Raw("UPDATE users SET block_status=false WHERE id=?",ID).Scan(&user)
	c.Redirect(303,"/admin/user_management/allusers")
}

//unblocking user
func UnBlockUser(c *gin.Context) {
	var user models.Users

	ID := c.Param("ID")
	if ID == "1"{
		dialog.Alert("Cant do anything,User has Admin privilages!")
		c.Redirect(303,"/admin/user_management/allusers")
		return
	}
	db := database.GetDb()
	db.Raw("SELECT block_status FROM users WHERE id=?",ID).Scan(&user)
	if user.Block_Status{
		dialog.Alert("Already in active!")
		c.Redirect(303,"/admin/user_management/allusers")
		return
	}

	db.Raw("UPDATE users SET block_status=true WHERE id=?",ID).Scan(&user)
	c.Redirect(303,"/admin/user_management/allusers")
	
}

//deleting user from database
func DeleteUser(c *gin.Context) {
	var user models.Users
	ID := c.Param("ID")
	if ID == "1"{
		dialog.Alert("Could not delete, User has Admin privilages!")
		c.Redirect(303,"/admin/user_management/allusers")
		return
	}
	db := database.GetDb()
	db.Raw("DELETE FROM users WHERE id=?",ID).Scan(&user)
	c.Redirect(303,"/admin/user_management/allusers")
	
}

//add user
func AddUsers(c *gin.Context) {
	
 c.HTML(200,"addusers.gohtml",nil)
}

func PAddUsers(c *gin.Context) {
	db := database.GetDb()
	FirstName := c.PostForm("firstname")
	LastName := c.PostForm("lastname")
	Email := c.PostForm("email")
	Password := c.PostForm("password")

	var users models.Users

	db.Raw("SELECT email FROM users WHERE email=?",Email).Scan(&users)
	if users.Email == Email{
		dialog.Alert("The Email is already taken!!")
		c.Redirect(303,"/admin/user_management/addusers")
		return
	}
	
		users.First_Name=FirstName
		users.Last_Name=LastName
		users.Email=Email
		users.Password=Password
		users.Block_Status=true
		
    db.Select("first_name","last_name","email","password","block_status").Create(&users)
	dialog.Alert("Account Create Successfully!")
	
	c.Redirect(303,"/admin/user_management/addusers")
}

