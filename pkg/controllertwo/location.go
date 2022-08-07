package controllertwo

import (
	"fmt"
	"log"

	"github.com/VJ-Vijay77/r4room/pkg/database"
	"github.com/VJ-Vijay77/r4room/pkg/models"
	"github.com/gin-gonic/gin"
)

func Location(c *gin.Context) {
	db := database.GetDb()
	session, err := Store.Get(c.Request, "session")
	if err != nil {
		log.Println("Cannot get sessions!")
	}
	useriD := session.Values["userID"]
	userID := fmt.Sprintf("%s", useriD)
	var userinfos models.Users

	db.Raw("SELECT first_name FROM users where email=?", userID).Scan(&userinfos)
	UserName := userinfos.First_Name

	var UserID int
	db.Raw("SELECT id FROM users WHERE email=?", userID).Scan(&UserID)
	//cart count
	var count int
	db.Raw("SELECT COUNT(user_id) FROM carts WHERE user_id=?", UserID).Scan(&count)
	//wishlist count
	var wishlistcount int
	db.Raw("SELECT COUNT(user_id) FROM wishlists WHERE user_id=?", UserID).Scan(&wishlistcount)

	location := c.Param("name")

	var lrooms []models.Rooms
	db.Where("location=?",location).Find(&lrooms)

	c.HTML(200,"locationrooms.gohtml",gin.H{
		"username": UserName,
		"count":    count,
		"wcount":   wishlistcount,
		"rooms":lrooms,
		"location":location,
	})

}
