package controller

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"strconv"
	"time"

	"github.com/VJ-Vijay77/r4room/pkg/database"
	"github.com/VJ-Vijay77/r4room/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//rendering admin homepage
func AdminHome(c *gin.Context) {

	//checking session
	ok := AdminLogedCheck(c)
	if !ok {
		c.Redirect(303, "/admin_login")
		return
	}
	session, _ := Store.Get(c.Request, "adminsession")
	email := session.Values["userID"]
	user_ID := fmt.Sprintf("%v", email)

	db := database.GetDb()

	var userinfos models.Users

	db.Raw("SELECT first_name,last_name FROM users where email=?", user_ID).Scan(&userinfos)
	UserName := userinfos.First_Name
	LastName := userinfos.Last_Name


	current := time.Now().Format("2006-01-02")

	// //daily sales
	// var dailysales []models.Orders
	// db.Where("checkindate=?", current).Find(&dailysales)
	//total daily revenue
	var TotalRevenue int64
	db.Raw("SELECT SUM(totalprice) FROM orders WHERE checkindate=?", current).Scan(&TotalRevenue)

	timed := time.Now()
	time2 := timed.AddDate(0, 0, -6)

	// weektimestart := timed.Format("2006-01-02")
	todaytime := time2.Format("2006-01-02")
	days1 := time2.Weekday()
	var weekdata []models.Orders
	db.Select("totalprice").Where("checkindate = ?", todaytime).Find(&weekdata)

	var first int
	for _, i := range weekdata {
		first += i.Totalprice
	}

	day2 := timed.AddDate(0, 0, -5)
	days2 := day2.Weekday()
	day2time := day2.Format("2006-01-02")
	var weekdata2 []models.Orders
	db.Select("totalprice").Where("checkindate = ?", day2time).Find(&weekdata2)
	var second int
	for _, i := range weekdata2 {
		second += i.Totalprice
	}

	day3 := timed.AddDate(0, 0, -4)
	days3 := day3.Weekday()

	day3time := day3.Format("2006-01-02")
	var weekdata3 []models.Orders
	db.Select("totalprice").Where("checkindate = ?", day3time).Find(&weekdata3)
	var third int
	for _, i := range weekdata3 {
		third += i.Totalprice
	}

	day4 := timed.AddDate(0, 0, -3)
	days4 := day4.Weekday()

	day4time := day4.Format("2006-01-02")
	var weekdata4 []models.Orders
	db.Select("totalprice").Where("checkindate = ?", day4time).Find(&weekdata4)
	var fourth int
	for _, i := range weekdata4 {
		fourth += i.Totalprice
	}

	day5 := timed.AddDate(0, 0, -2)
	days5 := day5.Weekday()

	day5time := day5.Format("2006-01-02")
	var weekdata5 []models.Orders
	db.Select("totalprice").Where("checkindate = ?", day5time).Find(&weekdata5)
	var five int
	for _, i := range weekdata5 {
		five += i.Totalprice
	}

	day6 := timed.AddDate(0, 0, -1)
	days6 := day6.Weekday()

	day6time := day6.Format("2006-01-02")
	var weekdata6 []models.Orders
	db.Select("totalprice").Where("checkindate = ?", day6time).Find(&weekdata6)
	var six int
	for _, i := range weekdata6 {
		six += i.Totalprice
	}

	day7 := time.Now()
	days7 := day7.Weekday()

	day7time := day7.Format("2006-01-02")
	var weekdata7 []models.Orders
	db.Select("totalprice").Where("checkindate = ?", day7time).Find(&weekdata7)
	var seven int
	for _, i := range weekdata7 {
		seven += i.Totalprice
	}

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

	//may stats starts here
	may := "2022-05-01"
	mayend := "2022-05-31"
	var maydetails []models.Orders
	var maycount int64
	db.Where("checkindate BETWEEN ? AND ?", may, mayend).Find(&maydetails)
	db.Model(&models.Orders{}).Where("checkindate BETWEEN ? AND ?", may, mayend).Count(&maycount)

	var maytotal int
	for _, i := range maydetails {
		maytotal += i.Totalprice
	}

	//yearly graph report
	yearstart := "2022-01-01"
	yearend := "2022-12-31"

	var yearlydetails []models.Orders
	var yearlycount int64
	db.Where("checkindate BETWEEN ? AND ?", yearstart, yearend).Find(&yearlydetails)
	db.Model(&models.Orders{}).Where("checkindate BETWEEN ? AND ?", yearstart, yearend).Count(&yearlycount)

	var yearlytotal int
	for _, i := range yearlydetails {
		yearlytotal += i.Totalprice
	}

	weektimestart := timed.Format("2006-01-02")
	todaytimes := time2.Format("2006-01-02")

	var weekdatas []models.Orders
	db.Where("checkindate > ?", todaytime).Find(&weekdata)

	var wrevenue int
	for _, i := range weekdata {
		wrevenue += i.Totalprice
	}

	c.HTML(200, "adminhome.gohtml", gin.H{
		"first":  first,
		"fday":   days1,
		"second": second,
		"sday":   days2,
		"third":  third,
		"tday":   days3,
		"fourth": fourth,
		"foday":  days4,

		"five":   five,
		"fifday": days5,

		"six":    six,
		"sixday": days6,

		"seven":  seven,
		"seday":  days7,
		"june":   monthlytotal,
		"may":    maytotal,
		"yearly": yearlytotal,
		"daily":  TotalRevenue,

		"wstart": weektimestart,
		"wend":   todaytimes,

		"weeklydetails": weekdatas,
		"wtotal":        wrevenue,
		"firstname":UserName,
		"lastname":LastName,
	})
}

func UserManagement(c *gin.Context) {

	//checking session
	ok := AdminLogedCheck(c)
	if !ok {
		c.Redirect(303, "/admin_login")
		return
	}

	c.HTML(200, "usermanagement.gohtml", nil)
}

//rendering all users page
func AllUsers(c *gin.Context) {

	//checking session
	ok := AdminLogedCheck(c)
	if !ok {
		c.Redirect(303, "/admin_login")
		return
	}

	db := database.GetDb()
	var data []models.Users

	//ordering all data by id
	db.Order("id").Find(&data)

	c.HTML(200, "allusers.gohtml", gin.H{
		"data": data,
	})
}

//rendering add room page
func AddRoom(c *gin.Context) {

	//checking session
	ok := AdminLogedCheck(c)
	if !ok {
		c.Redirect(303, "/admin_login")
		return
	}
	db := database.GetDb()
	var categories []models.Category
	db.Raw("SELECT id,category_name FROM category").Scan(&categories)

	c.HTML(200, "addroom.gohtml", gin.H{
		"category": categories,
	})
}

func PAddRoom(c *gin.Context) {
	var ok bool
	//checking session
	// ok := AdminLogedCheck(c)
	// if !ok {
	// 	c.Redirect(303, "/admin_login")
	// 	return
	// }
	db := database.GetDb()

	RoomName := c.PostForm("roomname")
	// CoverPicPath := c.PostForm("roompicpath")
	CoverPicPath, _ := c.FormFile("roompicpath")
	extension := filepath.Ext(CoverPicPath.Filename)
	Cover := uuid.New().String() + extension
	c.SaveUploadedFile(CoverPicPath, "./public/"+Cover)

	RoomPrice := c.PostForm("roomprice")
	RoomCost, _ := strconv.Atoi(RoomPrice)
	RoomCategory := c.PostForm("roomcategory")
	RoomDescription := c.PostForm("description")
	RoomLocation := c.PostForm("location")

	SubPicPath1, _ := c.FormFile("roompicpath1")
	extension = filepath.Ext(SubPicPath1.Filename)
	SubPic1 := uuid.New().String() + extension
	c.SaveUploadedFile(SubPicPath1, "./public/"+SubPic1)
	// SubPicPath1 := c.PostForm("roompicpath1")
	SubPicPath2, _ := c.FormFile("roompicpath2")
	extension = filepath.Ext(SubPicPath2.Filename)
	SubPic2 := uuid.New().String() + extension
	c.SaveUploadedFile(SubPicPath2, "./public/"+SubPic2)

	// SubPicPath2 := c.PostForm("roompicpath2")
	SubPicPath3, _ := c.FormFile("roompicpath3")
	extension = filepath.Ext(SubPicPath3.Filename)
	SubPic3 := uuid.New().String() + extension
	c.SaveUploadedFile(SubPicPath3, "./public/"+SubPic3)

	// SubPicPath3 := c.PostForm("roompicpath3")
	SubPicPath4, _ := c.FormFile("roompicpath4")
	extension = filepath.Ext(SubPicPath4.Filename)
	SubPic4 := uuid.New().String() + extension
	c.SaveUploadedFile(SubPicPath4, "./public/"+SubPic4)

	// SubPicPath4 := c.PostForm("roompicpath4")

	var addroom models.Rooms
	db.Raw("SELECT room_name FROM rooms WHERE room_name=?", RoomName).Scan(&addroom)
	if addroom.Room_Name == RoomName {
		ok = false
		k, _ := json.Marshal(ok)
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Write(k)

		return
	}

	addroom.Room_Name = RoomName
	addroom.Room_Price = RoomCost
	addroom.Category = RoomCategory
	addroom.Cover = Cover
	addroom.Sub1 = SubPic1
	addroom.Sub2 = SubPic2
	addroom.Sub3 = SubPic3
	addroom.Sub4 = SubPic4
	addroom.Status = "available"
	addroom.Description = RoomDescription
	addroom.Location = RoomLocation

	db.Select("room_name", "room_price", "category", "cover", "sub1", "sub2", "sub3", "sub4", "sub5", "status", "description", "location").Create(&addroom)

	ok = true
	k, _ := json.Marshal(ok)
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.Write(k)
	//c.Redirect(303, "/admin/add_room")
}

//rendering remove room page
func RemoveRoom(c *gin.Context) {

	//checking session
	ok := AdminLogedCheck(c)
	if !ok {
		c.Redirect(303, "/admin_login")
		return
	}
	db := database.GetDb()
	var room []models.Rooms
	db.Order("id").Find(&room)
	c.HTML(200, "removeroom.gohtml", gin.H{
		"allrooms": room,
	})
}

func DeleteRoom(c *gin.Context) {
	ID := c.Param("ID")

	db := database.GetDb()
	var rooms models.Rooms
	db.Raw("DELETE FROM rooms WHERE id=?", ID).Scan(&rooms)
	time.Sleep(1 * time.Second)
	c.Redirect(303, "/admin/remove_room")
}
