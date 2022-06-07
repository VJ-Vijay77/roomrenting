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
		c.Redirect(303, "/admin/user_management/allusers")
		return
	} 
	db.Raw("SELECT id,first_name,last_name,email,mobile FROM users WHERE first_name=?", Name).Scan(&results)

	var data []models.Users

	//ordering all data by id
	db.Order("id").Find(&data)

	c.HTML(200, "allusers.gohtml", gin.H{
		"userdata": results,
		"data":data,
	})
	//  c.Redirect(303, "/admin/user_management/addusers")
	// c.HTML(200,"addusers.gohtml",gin.H{"data":result,})
}
