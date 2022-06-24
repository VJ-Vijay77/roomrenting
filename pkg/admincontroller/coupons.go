package admincontroller

import (
	"github.com/VJ-Vijay77/r4room/pkg/database"
	"github.com/VJ-Vijay77/r4room/pkg/models"
	"github.com/gin-gonic/gin"
)

func Coupons(c *gin.Context){
	db := database.GetDb()

	var category []models.Category
	db.Raw("SELECT category_name FROM category").Scan(&category)
	c.HTML(200,"coupons.gohtml",gin.H{
		"category":category,
	})
}