package controller

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/VJ-Vijay77/r4room/pkg/database"
	"github.com/VJ-Vijay77/r4room/pkg/models"
	"github.com/gin-gonic/gin"
)

func Cart(c *gin.Context) {
	db := database.GetDb()
	var cartitems []models.Carts
	db.Find(&cartitems)
	c.HTML(200, "cart.gohtml", gin.H{
		"cart": cartitems,
	})
}

func PCart(c *gin.Context) {

	ok := UserLogedCheck(c)
	if !ok {
		var not = "login"
		k, _ := json.Marshal(not)
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Write(k)
		return

	}
	db := database.GetDb()
	session, _ := Store.Get(c.Request, "session")
	userID := session.Values["userID"]
	userEmail := fmt.Sprintf("%v", userID)

	I := c.Param("RID")
	RoomID, _ := strconv.Atoi(I)
	var user_ID int
	//  var roominfo models.Rooms
	var cart models.Carts
	db.Raw("SELECT id FROM users WHERE email=?", userEmail).Scan(&user_ID)
	var cartItems models.Carts
	 db.Where("room_id=? AND user_id=?",RoomID,user_ID).Find(&cartItems)
	if cartItems.Room_Id == RoomID && cartItems.User_Id==user_ID{
		var room = "sameroom"
		k, _ := json.Marshal(room)
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Write(k)
		return
	}
	cart.Room_Id = RoomID
	cart.User_Id = user_ID

	db.Select("room_id", "user_id").Create(&cart)
	var note = "added"
	k, _ := json.Marshal(note)
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.Write(k)
	

}
