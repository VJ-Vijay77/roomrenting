package controller

import "github.com/gin-gonic/gin"

func UserProfile(c *gin.Context){
	c.HTML(200,"userprofile.gohtml",nil)
}