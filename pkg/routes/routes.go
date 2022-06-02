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

	//Get methods
	api.GET("/user_login", controller.UserLogin)
	api.GET("/user_signup",controller.UserSignup)
	api.GET("/user_home",controller.UserHome)
	api.GET("/admin_login",controller.AdminLogin)
	api.GET("/admin_home",controller.AdminHome)
	
	//Post methods
	api.POST("/user_login", controller.PUserLogin)
	api.POST("/user_signup",controller.PUserSignup)
	api.POST("/admin_login",controller.PAdminLogin)


	//admin sub Get methods
	api.GET("/admin/user_management/allusers",controller.AllUsers)
	api.GET("/admin/user_management",controller.UserManagement)
	api.GET("/admin/add_room",controller.AddRoom)
	api.GET("/admin/remove_room",controller.RemoveRoom)

	//user management update delete add
	api.POST("/admin/user_management/fnupdate/:ID",controller.UpdateFUser)
	api.POST("/admin/user_management/lnupdate/:ID",controller.UpdateLUser)
	api.GET("/admin/user_management/block/:ID",controller.BlockUser)
	api.GET("/admin/user_management/unblock/:ID",controller.UnBlockUser)
	api.GET("/admin/user_management/delete/:ID",controller.DeleteUser)

	//add user 
	api.GET("/admin/user_management/addusers",controller.AddUsers)
	api.POST("/admin/user_management/addusers",controller.PAddUsers)

	//search user
	api.POST("/admin/user_management/searchusers",controller.SearchUsers)


}
