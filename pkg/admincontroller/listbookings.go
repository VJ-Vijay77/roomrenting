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

	//daily sales
	var dailysales []models.Orders
	db.Where("checkindate=?", current).Find(&dailysales)

	//total daily bookings
	var totaldailybookings int64
	//db.Raw("SELECT orderid FROM orders WHERE checkindate=?", current).Count(&totaldailybookings)
	db.Model(&models.Orderedrooms{}).Where("checkindate=?", current).Count(&totaldailybookings)

	//total daily revenue
	var TotalRevenue int64
	db.Raw("SELECT SUM(totalprice) FROM orders WHERE checkindate=?", current).Scan(&TotalRevenue)

	//total percentage of daily target
	var perc float64 = float64(TotalRevenue) / 50000 * 100
	dailypercentage := (int64(perc))

	//weekly stats
	var weeklydates []models.Orders
	db.Raw("SELECT firstname,checkindate,place,totalprice,roomnames FROM orders").Scan(&weeklydates)

	layout := "2006-01-02"
	var weeks time.Time
	var weekdates []time.Time

	for _, i := range weeklydates {
		weeks, _ = time.Parse(layout, i.Checkindate)
		weekdates = append(weekdates, weeks)
	}

	timed := time.Now()
	time2 := timed.AddDate(0, 0, -7)
	weektimestart := timed.Format("02-01-2006")
	todaytime := time2.Format("02-01-2006")
	var sevendate time.Time
	var sevendaysdate []string
	for _, i := range weekdates {
		if i.After(time2) {
			sevendate = i
			sevendaysdate = append(sevendaysdate, sevendate.Format("2006-01-02"))
		}
	}
	var wrevenue int
	var weeklybookedcount int
	var weekly models.Orders
	var weeklyd []models.Orders
	for _, o := range sevendaysdate {
		db.Where("checkindate=?", o).Find(&weekly)
		weeklyd = append(weeklyd, weekly)
		weeklybookedcount += 1
	}

	for _,i := range weeklyd{
		wrevenue += i.Totalprice
	}
	
	var floatperc float32 = float32(wrevenue)/400000*100
	weeklypercentage := (int64(floatperc))



	c.HTML(200, "sales.gohtml", gin.H{
		"dsales":    dailysales,
		"dtotal":    totaldailybookings,
		"drevenue":  TotalRevenue,
		"dailyperc": dailypercentage,

		"wstart":weektimestart,
		"wend":todaytime,
		"weeklycount":   weeklybookedcount,
		"weeklydetails": weeklyd,
		"wtotal" : wrevenue,
		"wperc" : weeklypercentage,
	})
}
