package controller

import (
	"fmt"

	"github.com/VJ-Vijay77/r4room/pkg/database"
	"github.com/VJ-Vijay77/r4room/pkg/models"
	"github.com/gin-gonic/gin"
)

//User home page
func UserHome(c *gin.Context) {

	//checking session
	ok := UserLogedCheck(c)
	if !ok {
		c.HTML(200,"userhome.gohtml",nil)
		return
	}
	db := database.GetDb()
	session, _ := Store.Get(c.Request, "session")
	useriD := session.Values["userID"]
	userID := fmt.Sprintf("%s", useriD)
	var userinfos models.Users
	db.Raw("SELECT first_name FROM users where email=?", userID).Scan(&userinfos)
	UserName := userinfos.First_Name
	c.HTML(200, "userhome.gohtml", gin.H{
		"data": userinfos,
		"username":UserName,
	})
}
