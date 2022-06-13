package admincontroller

import (
	"github.com/VJ-Vijay77/r4room/pkg/database"
	"github.com/VJ-Vijay77/r4room/pkg/models"
	"github.com/gin-gonic/gin"
)

func ListBookings(c *gin.Context){
	db := database.GetDb()

	var bookings []models.ListBookings
	db.Raw("SELECT orders.orderid,orderedrooms.roomid,orders.user_id,orders.totalprice,orderedrooms.status FROM orders INNER JOIN orderedrooms ON orders.orderid=orderedrooms.id").Scan(&bookings)


	c.HTML(200,"listbookings.gohtml",gin.H{
		"bookings":bookings,
	})
}