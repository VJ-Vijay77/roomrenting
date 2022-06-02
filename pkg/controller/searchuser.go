package controller

import (
	"github.com/VJ-Vijay77/r4room/pkg/database"
	"github.com/gin-gonic/gin"
	"tawesoft.co.uk/go/dialog"
)

//search user
func SearchUsers(c *gin.Context) {
	Name := c.PostForm("search")
	var result string
	db := database.GetDb()
	db.Raw("SELECT first_name FROM users WHERE first_name=?", Name).Scan(&result)
	if result == Name && result != ""{
		dialog.Alert("User Found")
	}else{
	dialog.Alert("User not found")
	}
	 c.Redirect(303, "/admin/user_management/addusers")
	// c.HTML(200,"addusers.gohtml",gin.H{"data":result,})
}
