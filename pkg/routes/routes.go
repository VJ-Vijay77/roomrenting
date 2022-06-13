package routes

import (
	"github.com/VJ-Vijay77/r4room/pkg/admincontroller"
	"github.com/VJ-Vijay77/r4room/pkg/controller"
	"github.com/VJ-Vijay77/r4room/pkg/controllertwo"
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
	api.GET("/user/room_booked_search",controller.BookedSearchRooms)
	api.GET("/user/room_info/:ID",controller.RoomInfo)

	//cart management
	api.GET("/user/cart",controller.Cart)
	api.GET("/user/cart/:RID",controller.AddToCart)
	api.GET("/user/cart/delete/:CID",controller.DeleteCart)


	//wishlist management
	api.GET("/user/wishlist",controller.Wishlish)
	api.GET("/user/wishlist/:RID",controller.AddToWishlist)
	api.GET("/user/wishlist/delete/:WID",controller.DeleteFromWishlist)
	api.GET("/user/wishlist/cart/:RID",controller.WLAddToCart)

	//payment gateway
	api.GET("/user/payment/:TotalPrice/:UserID",controller.Payment)
	api.POST("/user/payment/confirm",controller.PaymentConfirm)
	api.GET("/user/payment/success",controller.PaymentSuccess)

	//user profile
	api.GET("/user/profile",controller.UserProfile)
	api.GET("/user/order_history",controller.OrderHistory)
	api.GET("/user/edit_profile",controller.EditProfile)
	api.GET("/user/edit_profile_two",controller.EditProfileTwo)

	//user edit profile
	api.POST("/user/edit_profile/firstname/:ID",controller.EditProfileFirstName)
	api.POST("/user/edit_profile/lastname/:ID",controller.EditProfileLastName)
	api.POST("/user/edit_profile/mobile/:ID",controller.EditProfileMobile)
	api.POST("/user/edit_profile/email/:ID",controller.EditProfileEmail)
	api.POST("/user/edit_profile/editaddress/:ID",controller.EditProfileAddress)
	api.POST("/user/edit_profile/editaddresstwo/:ID",controller.AddressTwo)

	//settings
	api.GET("/user/settings",controllertwo.Settings)
	api.GET("/user/settings/change_password",controllertwo.ChangePassword)
	api.POST("/user/settings/change_password/:ID",controllertwo.PChangePassword)

	//user room checkings
	api.GET("/user/checkings",controllertwo.CheckingsList)
	api.GET("/user/rooms/checkedout/:RID/:UID",controllertwo.CheckingOut)

	//admin list bookings
	api.GET("/admin/list_bookings",admincontroller.ListBookings)
	
}
