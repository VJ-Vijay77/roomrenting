package admincontroller

import (
	"strconv"

	"github.com/VJ-Vijay77/r4room/pkg/database"
	"github.com/VJ-Vijay77/r4room/pkg/models"
	"github.com/gin-gonic/gin"
	"tawesoft.co.uk/go/dialog"
)

func Coupons(c *gin.Context) {
	db := database.GetDb()

	var category []models.Category
	db.Raw("SELECT category_name FROM category").Scan(&category)
	
	var coupons []models.Coupons
	db.Find(&coupons)

	c.HTML(200, "coupons.gohtml", gin.H{
		"category": category,
		"coup":coupons,
	})
}
func PCoupons(c *gin.Context) {
	db := database.GetDb()

	categoryname := c.PostForm("roomcategory")
	code := c.PostForm("couponcode")
	value := c.PostForm("couponvalue")
	Value, _ := strconv.Atoi(value)
	status := "true"

	db.AutoMigrate(models.Coupons{})
	var couponcheck []models.Coupons
	db.Find(&couponcheck)
	for _,i := range couponcheck{
		if i.Code == code{
			dialog.Alert("same coupon already exists")
	c.Redirect(303, "/admin/coupons")
			
			return 
		}
	}
	var coupons models.Coupons

	coupons.Category = categoryname
	coupons.Code = code
	coupons.Value = Value
	coupons.Status = status

	db.Select("category", "code", "value", "status").Create(&coupons)
	c.Redirect(303, "/admin/coupons")
}
