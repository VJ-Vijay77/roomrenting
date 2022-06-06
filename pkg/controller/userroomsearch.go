package controller

import (
	"github.com/VJ-Vijay77/r4room/pkg/database"
	"github.com/VJ-Vijay77/r4room/pkg/models"
	"github.com/gin-gonic/gin"
)


func SearchRooms(c *gin.Context) {
	db := database.GetDb()
	var rooms []models.Rooms
	db.Order("room_name").Find(&rooms)
	c.HTML(200,"roomsearchpage.gohtml",gin.H{
		"rooms":rooms,
	})
}