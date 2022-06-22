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

	var availablenow models.Rooms
	db.Raw("UPDATE rooms SET checkoutdate='availablenow' WHERE id=?",RID).Scan(&availablenow)

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
	// daily ends *** **** **** **** ends

	//weekly stats starts
	

	timed := time.Now()
	time2 := timed.AddDate(0, 0, -7)
	weektimestart := timed.Format("2006-01-02")
	todaytime := time2.Format("2006-01-02")
	
	var weekdata []models.Orders
	db.Where("checkindate > ?",todaytime).Find(&weekdata)
		var wcount int64
	db.Model(&models.Orders{}).Where("checkindate > ?", todaytime).Count(&wcount)


	var wrevenue int
	for _, i := range weekdata{
		wrevenue += i.Totalprice
	}

	var floatperc float32 = float32(wrevenue) / 400000 * 100
	weeklypercentage := (int64(floatperc))
	//weekly stats ends here ***** **** ****** ***//

	c.HTML(200, "sales.gohtml", gin.H{
		"dsales":    dailysales,
		"dtotal":    totaldailybookings,
		"drevenue":  TotalRevenue,
		"dailyperc": dailypercentage,

		"wstart":        weektimestart,
		"wend":          todaytime,
		"weeklycount":   wcount,
		"weeklydetails": weekdata,
		"wtotal":        wrevenue,
		"wperc":         weeklypercentage,
	})
}

func SalesReportMonthly(c *gin.Context) {
	db := database.GetDb()

	// june stats here
	k := "2022-05-31"
	var junedetails []models.Orders
	var junecount int64
	db.Where("checkindate > ?", k).Find(&junedetails)
	db.Model(&models.Orders{}).Where("checkindate > ?", k).Count(&junecount)

	var monthlytotal int
	for _, i := range junedetails {
		monthlytotal += i.Totalprice
	}

	var floatperc float32 = float32(monthlytotal) / 1000000 * 100
	monthlypercentage := (int64(floatperc))
	//june stats end

	//may stats starts here
	may := "2022-05-01"
	mayend := "2022-05-31"
	var maydetails []models.Orders
	var maycount int64
	db.Where("checkindate BETWEEN ? AND ?",may,mayend).Find(&maydetails)
	db.Model(&models.Orders{}).Where("checkindate BETWEEN ? AND ?",may,mayend).Count(&maycount)

	var maytotal int
	for _, i := range maydetails {
		maytotal += i.Totalprice
	}

	var mayperc float32 = float32(maytotal) / 1000000 * 100
	maypercentage := (int64(mayperc))
	// may end here
	c.HTML(200, "monthlysales.gohtml", gin.H{
		// june
		"junecount":   junecount,
		"juneperc":    monthlypercentage,
		"junedetails": junedetails,
		"junetotal":   monthlytotal,
		// may
		"maycount":   maycount,
		"mayperc":    maypercentage,
		"maydetails": maydetails,
		"maytotal":   maytotal,

	})
}

func SalesReportYearly(c *gin.Context) {
	db := database.GetDb()
	yearstart := "2022-01-01"
	yearend := "2022-12-31"

	var yearlydetails []models.Orders
	var yearlycount int64
	db.Where("checkindate BETWEEN ? AND ?",yearstart,yearend).Find(&yearlydetails)
	db.Model(&models.Orders{}).Where("checkindate BETWEEN ? AND ?",yearstart,yearend).Count(&yearlycount)

	var yearlytotal int
	for _, i := range yearlydetails {
		yearlytotal += i.Totalprice
	}

	var yearlyperc float32 = float32(yearlytotal) / 10000000 * 100
	yearlypercentage := (int64(yearlyperc))

	c.HTML(200, "yearlysales.gohtml", gin.H{
		"yearcount":   yearlycount,
		"yearperc":    yearlypercentage,
		"yeardetails": yearlydetails,
		"yeartotal":   yearlytotal,
	})
}

