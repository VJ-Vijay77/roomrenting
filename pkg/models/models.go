package models

//model representing the table in database
type Users struct {
	ID           uint
	First_Name   string
	Last_Name    string
	Email        string
	Password     string
	Role         string
	Block_Status bool
	Mobile       string
}

type Rooms struct {
	ID         int
	Room_Name  string
	Room_Price int
	Category   string
	Cover      string
	Sub1       string
	Sub2       string
	Sub3       string
	Sub4       string
	Sub5       string
	Status     string
	Checkoutdate string
	Description string
}

type Category struct {
	ID            int
	Category_Name string
}

type Carts struct {
	Cartsid     int
	Cartsroomid int
	User_Id     int
	Days        int
	Total       int
	Startdate   string
	Endate      string
}

type Cart_Infos struct {
	Cartsid     int
	Room_Name   string
	Room_Price  string
	Cartsroomid int
	Cover       string
	Category    string
	User_Id     int
	User_Name   string
	Days        int
	Total       int
	Startdate   string
	Endate      string
	
}

type Wishlists struct {
	Wishlistid  int
	User_ID     int
	Wishroomsid int
}

type Wishlist_Infos struct {
	Wishlistid  int
	Room_Name   string
	Room_Price  string
	Wishroomsid int
	Cover       string
	Category    string
	User_Id     int
	Status      string
}

type Orders struct {
	Orderid       int
	User_ID       int
	Firstname     string
	Lastname      string
	Housename     string
	Place         string
	State         string
	Mobile        string
	Totalprice    int
	Roomnames     string
	Paymentmethod string
	Accountholder string
	Checkindate string
	Checkoutdate string
}

type Orderedrooms struct {
	ID      int
	Roomid  int
	User_Id int
	Status  string
	Checkindate string
	Checkoutdate string
}

type Useraddress struct {
	Adrid     int
	User_Id   int
	Housename string
	Place     string
	Mobile    string
	State     string
	PIN       int
	Main      string
}

type BookingDetails struct {
	User_Id    int
	Id         int
	Room_Name  string
	Status     string
	Cover      string
	Category   string
	Room_Price string
}

type ListBookings struct {
	Orderid    int
	Roomid     int
	User_Id    int
	Totalprice string
	Status     string
}

type Razorpaydetails struct {
	User_Id   int
	Payid     string
	Orderid   string
	Signature string
}

type TotalPrice struct {
	Total string
	Days  int
}

type Dates struct{
	User_Id int
	Cartsroomid int
	Startdate string
	Endate string
}

