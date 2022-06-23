package controllertwo

import (
	"fmt"
	"strconv"
	"time"

	"github.com/VJ-Vijay77/r4room/pkg/database"
	"github.com/VJ-Vijay77/r4room/pkg/models"
	"github.com/gin-gonic/gin"
)

func CheckingsList(c *gin.Context) {
	db := database.GetDb()
	session, _ := Store.Get(c.Request, "session")
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

	var bookings []models.BookingDetails
	
db.Raw("SELECT user_id,rooms.id,rooms.room_name,orderedrooms.status,rooms.cover,rooms.category,rooms.room_price,rooms.offers,rooms.value,rooms.discountprice,rooms.description FROM orderedrooms INNER JOIN rooms ON rooms.id=orderedrooms.roomid WHERE user_id=?",UserID).Scan(&bookings)
	
	
	
	c.HTML(200, "checkingslist.gohtml", gin.H{
		"data":     userinfos,
		"username": UserName,
		"count":    count,
		"wcount":   wishlistcount,
		"UID":UserID,
		"bookings":bookings,
	})
}



func CheckingOut(c *gin.Context){
	RoomId := c.Param("RID")
	RID,_ := strconv.Atoi(RoomId)
	userId := c.Param("UID")
	UID,_ := strconv.Atoi(userId)
	db := database.GetDb()

	var roomupdations models.Orderedrooms
	db.Raw("UPDATE orderedrooms SET status='checkedout' WHERE roomid=? AND user_id=?",RID,UID).Scan(&roomupdations)

	var roomstatus models.Rooms
	db.Raw("UPDATE rooms SET status='available' WHERE id=?",RID).Scan(&roomstatus)

	var availablenow models.Rooms
	db.Raw("UPDATE rooms SET checkoutdate='availablenow' WHERE id=?",RID).Scan(&availablenow)
	time.Sleep(2 *time.Second)

	c.Redirect(303,"/user/checkings")
}

func Cancel (c *gin.Context){
	db := database.GetDb()
	session, _ := Store.Get(c.Request, "session")
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


	RoomId := c.Param("RID")
	RID,_ := strconv.Atoi(RoomId)
	userId := c.Param("UID")
	UID,_ := strconv.Atoi(userId)
	Price := c.Param("Total")
	Total,_ := strconv.Atoi(Price)
	
	
	c.HTML(200,"cancelbookings.gohtml",gin.H{
		"data":     userinfos,
		"username": UserName,
		"count":    count,
		"wcount":   wishlistcount,
		"rid":RID,
		"uid":UID,
		"total":Total,
		"refund":Total/2,
	})

}


func Refund(c *gin.Context){
	db := database.GetDb()
	RoomId := c.Param("RID")
	RID,_ := strconv.Atoi(RoomId)
	userId := c.Param("UID")
	UID,_ := strconv.Atoi(userId)
	Price := c.Param("Total")
	Total,_ := strconv.Atoi(Price)
	  fmt.Println(db,RID,UID,Total)
	db.AutoMigrate(&models.Wallets{})
	var wallet []models.Wallets
	db.Find(&wallet)

	var newallet models.Wallets
	var check bool
	for _,i := range wallet{
		if i.User_Id == UID{
			check = true
			db.Where("user_id=?",UID).Find(&newallet)
		}
	}
	
	var balance int = newallet.Balance+Total
	 
	if check {
		db.Raw("UPDATE wallets SET balance=? WHERE user_id=?",balance,UID).Scan(&newallet)
		
	}

	var create models.Wallets
	if !check{
		create.User_Id=UID
		create.Balance=Total
		db.Select("user_id","balance").Create(&create)
	}


	var roomupdations models.Orderedrooms
	db.Raw("UPDATE orderedrooms SET status='checkedout' WHERE roomid=? AND user_id=?",RID,UID).Scan(&roomupdations)

	var roomstatus models.Rooms
	db.Raw("UPDATE rooms SET status='available' WHERE id=?",RID).Scan(&roomstatus)

	var availablenow models.Rooms
	db.Raw("UPDATE rooms SET checkoutdate='availablenow' WHERE id=?",RID).Scan(&availablenow)

	c.Redirect(303,"/user/refund/success")	
	
}
func RSuccess (c *gin.Context){
	c.HTML(200,"refundsuccess.gohtml",nil)
}