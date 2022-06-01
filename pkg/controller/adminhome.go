package controller

import (
	"github.com/gin-gonic/gin"
)

func AdminHome(c *gin.Context) {
	c.HTML(200, "adminhome.gohtml", nil)
}
