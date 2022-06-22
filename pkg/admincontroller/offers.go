package admincontroller

import (
	"strconv"

	"github.com/VJ-Vijay77/r4room/pkg/database"
	"github.com/VJ-Vijay77/r4room/pkg/models"
	"github.com/gin-gonic/gin"
	"tawesoft.co.uk/go/dialog"
)

func Offers(c *gin.Context) {
	db := database.GetDb()
	var category []models.Category
	db.Raw("SELECT category_name FROM category").Scan(&category)
	
	var rooms []models.Rooms
	db.Distinct("category","offers","value").Find(&rooms)
	
	c.HTML(200, "offers.gohtml", gin.H{
		"name": category,
		"rooms":rooms,
	})
}

func CategoryOffers(c *gin.Context) {
	db := database.GetDb()

	name := c.PostForm("roomcategory")
	offerpercenteage := c.PostForm("value")
	perc,_ := strconv.Atoi(offerpercenteage)
	var update []models.Rooms
	db.Raw("UPDATE rooms SET value=? WHERE category=?", perc, name).Scan(&update)
	db.Raw("UPDATE rooms SET offers='true' WHERE category=?",name).Scan(&update)
	dialog.Alert("Offer Updated Successfully!!")
	c.Redirect(303, "/admin/offers")
}

func StopOffers(c *gin.Context){
	db := database.GetDb()
	cname := c.Param("name")

	var stop []models.Rooms
	db.Raw("UPDATE rooms SET offers='false' WHERE category=?",cname).Scan(&stop)

	c.Redirect(303,"/admin/offers")
}