package controller

import (
	"github.com/VJ-Vijay77/r4room/pkg/database"
	"github.com/gin-gonic/gin"
)

func Payment(c *gin.Context) {
	TotalPrice := c.Param("TotalPrice")
	UserID := c.Param("UserID")

	db := database.GetDb()
	//taking order details
	var roomnames []string
	db.Raw("SELECT rooms.room_name FROM carts INNER JOIN rooms ON rooms.id=carts.cartsroomid WHERE carts.user_id=?", UserID).Scan(&roomnames)

	c.HTML(200, "payment.gohtml", gin.H{
		"total":     TotalPrice,
		"roomnames": roomnames,
	})
}
