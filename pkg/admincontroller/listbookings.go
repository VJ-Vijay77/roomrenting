package admincontroller

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/VJ-Vijay77/r4room/pkg/database"
	"github.com/VJ-Vijay77/r4room/pkg/models"
	"github.com/gin-gonic/gin"
)

func ListBookings(c *gin.Context) {
	db := database.GetDb()

	var bookings []models.ListBookings
	db.Raw("SELECT orders.orderid,orderedrooms.roomid,orders.user_id,orders.totalprice,orderedrooms.status FROM orders INNER JOIN orderedrooms ON orders.orderid=orderedrooms.id").Scan(&bookings)

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

func SalesReport(c *gin.Context) {
	db := database.GetDb()
	current := time.Now().Format("2006-01-02")

	var dailysales []models.Orders
	db.Where("checkindate=?", current).Find(&dailysales)

	var totaldailybookings int64
	db.Raw("SELECT orderid FROM orders WHERE checkindate=?", current).Count(&totaldailybookings)

	var TotalRevenue int64
	db.Raw("SELECT SUM(totalprice) FROM orders WHERE checkindate=?", current).Scan(&TotalRevenue)
	
	var perc float64 = float64(TotalRevenue) / 100000 * 100

	dailypercentage := (int64(perc))
	c.HTML(200, "sales.gohtml", gin.H{
		"dsales":    dailysales,
		"dtotal":    totaldailybookings,
		"drevenue":  TotalRevenue,
		"dailyperc": dailypercentage,
		 
	})
}
