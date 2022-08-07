package controllertwo

import (
	"encoding/json"
	"fmt"
	"path/filepath"

	"github.com/VJ-Vijay77/r4room/pkg/database"
	"github.com/VJ-Vijay77/r4room/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Wallet(c *gin.Context) {
	db := database.GetDb()
	var wallet models.Wallets
	db.Select("balance").Find(&wallet)

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

	c.HTML(200, "wallet.gohtml", gin.H{
		"bal":      wallet.Balance,
		"username": UserName,
		"count":    count,
		"wcount":   wishlistcount,
	})
}

func ProfilePic(c *gin.Context) {
	db := database.GetDb()
	session, _ := Store.Get(c.Request, "session")
	useriD := session.Values["userID"]
	userID := fmt.Sprintf("%s", useriD)

	var UserID int
	db.Raw("SELECT id FROM users WHERE email=?", userID).Scan(&UserID)

	var ok string

	Profile, _ := c.FormFile("pfpic")
	if Profile.Filename == "" {
		ok = "nil"
		k, _ := json.Marshal(ok)
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Write(k)
		return

	}
	extension := filepath.Ext(Profile.Filename)
	ProfiePicName := uuid.New().String() + extension
	c.SaveUploadedFile(Profile, "./public/profile/"+ProfiePicName)

	db.AutoMigrate(&models.Profilepics{})

	var cprof []models.Profilepics
	db.Find(&cprof)
	var checkppic models.Profilepics
	for _,i := range cprof{
		if i.User_Id == UserID{
			db.Raw("UPDATE profilepics SET pic=? WHERE user_id=?",ProfiePicName,UserID).Scan(&checkppic)
			ok = "updated"
		k, _ := json.Marshal(ok)
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Write(k)
			return
		}
	}
	
	var profilepics models.Profilepics
	profilepics.Pic = ProfiePicName
	profilepics.User_Id = UserID
	db.Create(&profilepics)
	    ok = "added"
		k, _ := json.Marshal(ok)
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Write(k)
		

}
