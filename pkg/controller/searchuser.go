package controller

import (
	"github.com/VJ-Vijay77/r4room/pkg/database"
	"github.com/VJ-Vijay77/r4room/pkg/models"
	"github.com/gin-gonic/gin"
	"tawesoft.co.uk/go/dialog"
)

//search user
func SearchUsers(c *gin.Context) {
	Name := c.PostForm("search")
	var result string
	var results []models.Users
	db := database.GetDb()

	db.Raw("SELECT first_name FROM users WHERE first_name=?", Name).Scan(&result)

	if result != Name || result == "" {
		dialog.Alert("User Not Found")
		c.Redirect(303, "/admin/user_management/addusers")
		return
	} 
	db.Raw("SELECT id,first_name,last_name,email,mobile FROM users WHERE first_name=?", Name).Scan(&results)

	c.HTML(200, "addusers.gohtml", gin.H{
		"userdata": results,
	})
	//  c.Redirect(303, "/admin/user_management/addusers")
	// c.HTML(200,"addusers.gohtml",gin.H{"data":result,})
}
