package controller

import (
	"github.com/gin-gonic/gin"
)

//User home page
func UserHome(c *gin.Context) {

	c.HTML(200, "userhome.gohtml", gin.H{})
}
