package admincontroller

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/VJ-Vijay77/r4room/pkg/database"
	"github.com/VJ-Vijay77/r4room/pkg/models"
	"github.com/gin-gonic/gin"
)

func ListBookings(c *gin.Context) {
	db := database.GetDb()

	var bookings []models.ListBookings
	db.Raw("SELECT orders.orderid,orderedrooms.roomid,orders.user_id,orders.totalprice,orderedrooms.status FROM orders INNER JOIN orderedrooms ON orders.orderid=orderedrooms.id").Scan(&bookings)
	
	
	fmt.Println(bookings)
	c.HTML(200, "listbookings.gohtml", gin.H{
		"bookings": bookings,
	})
}

func CheckoutBookings(c *gin.Context) {
	var ok string
	db := database.GetDb()
	RoomID := c.Param("RID")
	OrderID := c.Param("OID")

	RID, _ := strconv.Atoi(RoomID)
	OID, _ := strconv.Atoi(OrderID)

	var roomsupdation models.Rooms
	db.Raw("UPDATE rooms SET status='available' WHERE id=?", RID).Scan(&roomsupdation)

	var orderedupdation models.Orderedrooms
	db.Raw("UPDATE orderedrooms SET status='checkedout' WHERE id=?", OID).Scan(&orderedupdation)

	ok = "true"
	k, _ := json.Marshal(ok)
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.Write(k)

}
