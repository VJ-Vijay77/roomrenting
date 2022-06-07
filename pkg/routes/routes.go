package routes

import (
	"github.com/VJ-Vijay77/r4room/pkg/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(api *gin.Engine) {
	//loading gohtml files from templates directory
	api.LoadHTMLGlob("templates/*.gohtml")
	
	//loading static files like css or js
	api.Static("static/","./static")

	//index router
	api.GET("/",controller.Index)

	//Get methods
	api.GET("/user_login", controller.UserLogin)
	api.GET("/login_with_otp", controller.UserOtpLogin)
	api.GET("/validate_otp", controller.ValidateOtp)
	api.GET("/user_signup",controller.UserSignup)
	api.GET("/user_home",controller.UserHome)
	api.GET("/admin_login",controller.AdminLogin)
	api.GET("/admin_home",controller.AdminHome)
	
	//Post methods
	api.POST("/user_login", controller.PUserLogin)
	api.POST("/login_with_otp", controller.PUserOtpLogin)
	api.POST("/validate_otp", controller.PValidateOtp)
	api.POST("/user_signup",controller.PUserSignup)
	api.POST("/admin_login",controller.PAdminLogin)


	//admin sub Get methods
	api.GET("/admin/user_management/allusers",controller.AllUsers)
	api.GET("/admin/user_management",controller.UserManagement)
	api.GET("/admin/add_room",controller.AddRoom)
	api.POST("/admin/add_room",controller.PAddRoom)
	api.GET("/admin/remove_room",controller.RemoveRoom)
	api.GET("/admin/remove_room/delete/:ID",controller.DeleteRoom)

	//user management update delete add
	api.POST("/admin/user_management/fnupdate/:ID",controller.UpdateFUser)
	api.POST("/admin/user_management/lnupdate/:ID",controller.UpdateLUser)
	api.Any("/admin/user_management/block/:ID",controller.BlockUser)
	api.GET("/admin/user_management/unblock/:ID",controller.UnBlockUser)
	api.GET("/admin/user_management/delete/:ID",controller.DeleteUser)

	//add user 
	api.GET("/admin/user_management/addusers",controller.AddUsers)
	api.POST("/admin/user_management/addusers",controller.PAddUsers)

	//search user
	api.POST("/admin/user_management/searchusers",controller.SearchUsers)

	//user/admin logout 
	api.GET("/user_logout",controller.UserLogout) 
	api.GET("/admin_logout",controller.AdminLogout) 


	//room searching and info pages
	api.GET("/user/room_search",controller.SearchRooms)
	api.GET("/user/room_info/:ID",controller.RoomInfo)


}
