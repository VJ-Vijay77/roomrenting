package controller

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/VJ-Vijay77/r4room/pkg/database"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"

	"github.com/VJ-Vijay77/r4room/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"tawesoft.co.uk/go/dialog"
)

//setting sessions using gorilla sesions
var Store = sessions.NewCookieStore([]byte(os.Getenv("session")))

//User login get and post methods
func UserLogin(c *gin.Context) {

	//checking session
	ok := UserLogedCheck(c)
	if ok {
		c.Redirect(303, "/user_home")
		return
	}

	c.HTML(200, "userlogin.gohtml", nil)
}

//validating user credentials from database
func PUserLogin(c *gin.Context) {
	db := database.GetDb()
	Email := c.PostForm("useremail")
	Password := c.PostForm("userpassword")

	var ds models.Users
	db.Raw("SELECT email,password,block_status FROM users WHERE email=? AND password=?", Email, Password).Scan(&ds)
	if Email != ds.Email || Password != ds.Password {
		dialog.Alert("Username or Password is Incorrect\n\t\tTry again!")
		c.Redirect(303, "/user_login")
		return
	}
	if !ds.Block_Status {
		dialog.Alert("You have been blocked by Admin!\n  contact:vijayadmin@admin")
		c.Redirect(303, "/user_login")
		return
	}

	//creating sessions for the user
	session, err := Store.Get(c.Request, "session")
	if err != nil {
		log.Println("Error getting session !!")
	}
	session.Values["userID"] = Email
	session.Save(c.Request, c.Writer)
	c.Redirect(303, "/user_home")

}

//user otp login
func UserOtpLogin(c *gin.Context) {
	c.HTML(200, "otppage.gohtml", nil)
}

func PUserOtpLogin(c *gin.Context) {
	Mobile := c.PostForm("mobile")
	db := database.GetDb()
	var mob models.Users
	db.Raw("SELECT mobile FROM users WHERE mobile=?", Mobile).Scan(&mob)

	if mob.Mobile != Mobile {
		dialog.Alert("Mobile number doesnt exist! Create new one.")
		c.Redirect(303, "/login_with_otp")
		return
	}

	store,_ := Store.Get(c.Request,"session")
	store.Values["userMob"]=Mobile
	store.Save(c.Request,c.Writer)
	c.Redirect(303,"/validate_otp") 
	


}

func ValidateOtp(c *gin.Context) {
	store,_ := Store.Get(c.Request,"session")
	Mob := store.Values["userMob"]
	Converted := fmt.Sprintf("%v",Mob)
	Mobile := "+91"+Converted
	
	rand.Seed(time.Now().UnixNano())

	value := rand.Intn(9999-1000) + 1000
	otp := strconv.Itoa(value)

	accountSid := "ACdac52be2b71187a48f36876f6343f734"
	authToken := "7cdff407c08ba5cb0ed3e13fbe255776"

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})
	
	params := &openapi.CreateMessageParams{}
	params.SetTo(Mobile)
	params.SetFrom("+19844647150")
	params.SetBody("Hello,Your OTP for logging into R4Rooms is - "+otp)

	_, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println(err.Error())
		err = nil
	}else{
		fmt.Println("OTP sent successfully!")
	}
	

	
	store.Values["OTP"]=otp
	store.Save(c.Request,c.Writer)


	c.HTML(200, "validateotp.gohtml", nil)

}


func PValidateOtp(c *gin.Context) {
	store,_ := Store.Get(c.Request,"session")
	 OTPcode:= store.Values["OTP"]
	 Mob := store.Values["userMob"]
	 Mobile := fmt.Sprintf("%v",Mob)
	OTP := fmt.Sprintf("%v",OTPcode)
	userOTP := c.PostForm("otp")
	db := database.GetDb()

	if OTP != userOTP {
		store.Options = &sessions.Options{
			Path:     "/",
			MaxAge:   -1,
			HttpOnly: true,
		}
		sessions.Save(c.Request, c.Writer)
		dialog.Alert("OTP is Incorrect !")
		c.Redirect(303,"/login_with_otp")
		return
	}
	dialog.Alert("OTP verified successfully!")
	
	var userinfo models.Users
	db.Raw("SELECT first_name,email FROM users WHERE mobile=?",Mobile).Scan(&userinfo)
	store.Values["userID"]=userinfo.Email
	store.Save(c.Request,c.Writer)
	dialog.Alert("Welcome %s !!",userinfo.First_Name)


	c.Redirect(303,"/user_home")


	
}





func UserLogout(c *gin.Context) {
	session, err := Store.Get(c.Request, "session")
	if err != nil {
		log.Println("couldnt fetch sessions!")
	}
	session.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	}
	sessions.Save(c.Request, c.Writer)
	c.Redirect(303, "/user_login")
}
