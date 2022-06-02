package controller

import (
	"github.com/gin-gonic/gin"
)

//User home page
func UserHome(c *gin.Context) {

	//checking session
	ok := UserLogedCheck(c)
	if !ok {
		c.Redirect(303, "/user_login")
		return
	}

	c.HTML(200, "userhome.gohtml", gin.H{})
}
