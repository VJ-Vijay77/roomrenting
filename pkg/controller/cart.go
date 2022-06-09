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
	session,_ := Store.Get(c.Request,"session")
	email := session.Values["userID"]
	user_ID := fmt.Sprintf("%v",email)

	db := database.GetDb()
	var UserID int
	db.Raw("SELECT id FROM users WHERE email=?",user_ID).Scan(&UserID)

	var userinfos models.Users
	db.Raw("SELECT first_name FROM users where email=?", user_ID).Scan(&userinfos)
	UserName := userinfos.First_Name
	var cartitems []models.Cart_Infos
	db.Raw("SELECT carts.cartsid,users.id,users.first_name,rooms.room_name,carts.cartsroomid,rooms.cover,rooms.room_price,rooms.category FROM carts INNER JOIN rooms ON rooms.id=carts.cartsroomid INNER JOIN users ON carts.user_id=users.id WHERE user_id=?",UserID).Scan(&cartitems)
	c.HTML(200, "cart.gohtml", gin.H{
		"cart": cartitems,
		"username":UserName,
	})
}

func AddToCart(c *gin.Context) {

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
	 db.Where("cartsroomid=? AND user_id=?",RoomID,user_ID).Find(&cartItems)
	if cartItems.Cartsroomid == RoomID && cartItems.User_Id==user_ID{
		var room = "sameroom"
		k, _ := json.Marshal(room)
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Write(k)
		return
	}
	cart.Cartsroomid = RoomID
	cart.User_Id = user_ID

	db.Select("cartsroomid", "user_id").Create(&cart)
	var note = "added"
	k, _ := json.Marshal(note)
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.Write(k)
	

}



// func DeleteCart (c *gin.Context){
// 	CID := c.Param("CID")
// 	CartId,_ := strconv.Atoi(CID)
// 	db := database.GetDb()
// 	var cart models.Carts
// 	db.Raw("")

// }