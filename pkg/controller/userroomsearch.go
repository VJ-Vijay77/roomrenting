package controller

import "github.com/gin-gonic/gin"


func SearchRooms(c *gin.Context) {
	c.HTML(200,"roomsearchpage.gohtml",nil)
}