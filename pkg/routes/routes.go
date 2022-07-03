package routes

import (
	"github.com/VJ-Vijay77/r4room/pkg/admincontroller"
	"github.com/VJ-Vijay77/r4room/pkg/controller"
	"github.com/VJ-Vijay77/r4room/pkg/controllertwo"
	"github.com/gin-gonic/gin"
)

func InitRoutes(api *gin.Engine) {
	//loading gohtml files from templates directory
	api.LoadHTMLGlob("templates/**/*.gohtml")
	
	//loading images for rooms
	api.Static("public/", "./public")
	
	//loading static files like css or js
	api.Static("static/","./static")

	//index router
	api.GET("/",controller.Index)

	//login Get methods
	api.GET("/user_login", controller.UserLogin)
	api.GET("/login_with_otp", controller.UserOtpLogin)
	api.GET("/validate_otp", controller.ValidateOtp)
	api.GET("/user_signup",controller.UserSignup)
	api.GET("/user_home",controller.UserHome)
	api.GET("/admin_login",controller.AdminLogin)
	api.GET("/admin_home",controller.AdminHome)
	
	//login Post methods
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
	api.GET("/user/room_search",controller.AvailableSearchRooms)
	api.GET("/user/all_room_search",controller.AllSearchRooms)
	api.GET("/user/all_room_search/filter/:start",controller.DateFilter)
	api.GET("/user/room_booked_search",controller.BookedSearchRooms)
	api.GET("/user/room_info/:ID",controller.RoomInfo)

	//filter by category
	api.GET("/user/filter/singleroom",controllertwo.SingleRoomFilter)
	api.GET("/user/filter/avsingleroom",controllertwo.AVSingleRoomFilter)
	api.GET("/user/filter/bksingleroom",controllertwo.BKSingleRoomFilter)
	api.GET("/user/filter/doubleroom",controllertwo.DoubleRoomFilter)
	api.GET("/user/filter/avdoubleroom",controllertwo.AVDoubleRoomFilter)
	api.GET("/user/filter/bkdoubleroom",controllertwo.BKDoubleRoomFilter)

	//filter by prize
	api.GET("/user/filter/abovefive",controllertwo.AboveFive)
	api.GET("/user/filter/aboveseven",controllertwo.AboveSeven)
	api.GET("/user/filter/belowfive",controllertwo.BelowFive)
	api.GET("/user/filter/avabovefive",controllertwo.AVAboveFive)
	api.GET("/user/filter/avaboveseven",controllertwo.AVAboveSeven)
	api.GET("/user/filter/avbelowfive",controllertwo.AVBelowFive)
	api.GET("/user/filter/bkabovefive",controllertwo.BKAboveFive)
	api.GET("/user/filter/bkaboveseven",controllertwo.BKAboveSeven)
	api.GET("/user/filter/bkbelowfive",controllertwo.BKBelowFive)


	//cart management
	api.GET("/user/cart",controller.Cart)
	api.GET("/user/cart/:RID/:Days/:Startdate/:Endate",controller.AddToCart)
	api.GET("/user/cart/delete/:CID",controller.DeleteCart)


	//wishlist management
	api.GET("/user/wishlist",controller.Wishlish)
	api.GET("/user/wishlist/:RID",controller.AddToWishlist)
	api.GET("/user/wishlist/delete/:WID",controller.DeleteFromWishlist)
	api.GET("/user/wishlist/cart/:RID",controller.WLAddToCart)

	//payment gateway
	api.GET("/user/payment/:TotalPrice/:UserID/:Roomid/:Startdate/:Endate",controller.Payment)
	api.POST("/user/payment/addresspick/:TotalPrice/:UserID/:Roomid/:Startdate/:Endate",controller.PaymentAddressPick)
	//payment at office
	api.POST("/user/payment/confirm/:start/:end/:wallet/:total",controller.PaymentConfirm)
	api.GET("/user/payment/success",controller.PaymentSuccess)
	//payment Rozor pay
	api.GET("/user/payment/razorpay/:total/:uid/:start/:end/:wallet/:roomid",controller.RazorPay)
	api.GET("/user/payment/razorypaysuccess/:rpid/:roid/:rsign/:start/:end/:total/:wallet/:roomid",controller.RazorPaySuccess)
	api.GET("/user/payment/razorpay/successpage/:pid/:orderid/:signature/:refid",controller.RPSuccess)

	//user profile
	api.GET("/user/profile",controller.UserProfile)
	api.GET("/user/order_history",controller.OrderHistory)
	api.GET("/user/edit_profile",controller.EditProfile)
	api.GET("/user/edit_profile_two",controller.EditProfileTwo)
	api.GET("/user/profile/wallet",controllertwo.Wallet)
	api.POST("/user/profile/profilepic",controllertwo.ProfilePic)

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
	api.GET("/user/rooms/cancel/:RID/:UID/:Total",controllertwo.Cancel)
	api.POST("/user/refund/:RID/:UID/:Total",controllertwo.Refund)
	//admin list bookings
	api.GET("/admin/list_bookings",admincontroller.ListBookings)
	api.GET("/admin/list_bookings/checkout/:RID/:OID",admincontroller.CheckoutBookings)

	//Admin sales report
	api.GET("/admin/salesreport",admincontroller.SalesReport)
	api.GET("/admin/salesreport/monthly",admincontroller.SalesReportMonthly)
	api.GET("/admin/salesreport/yearly",admincontroller.SalesReportYearly)
	api.GET("/admin/salesreport/graph",admincontroller.SalesReportGraph)

	//offers
	api.GET("/admin/offers",admincontroller.Offers)
	api.POST("/admin/offers/update",admincontroller.CategoryOffers)
	api.GET("/admin/offers/stop/:name",admincontroller.StopOffers)
	api.GET("/admin/coupons",admincontroller.Coupons)
	api.POST("/admin/coupons",admincontroller.PCoupons)
	api.GET("/admin/deletecoupons/:ID",admincontroller.DeleteCoupons)
	api.GET("/admin/inactivecoupons/:ID",admincontroller.InactiveCoupons)
	api.GET("/admin/activecoupons/:ID",admincontroller.ActiveCoupons)


	//refund routes
	api.GET("/user/refund/success",controllertwo.RSuccess)


}
