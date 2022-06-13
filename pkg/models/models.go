package models

//model representing the table in database
type Users struct{
	ID uint 
	First_Name string 
	Last_Name string 
	Email string 
	Password string 
	Role string
	Block_Status bool
	Mobile string
}

type Rooms struct{
	ID int
	Room_Name string
	Room_Price string
	Category string
	Cover string
	Sub1 string
	Sub2 string
	Sub3 string
	Sub4 string
	Sub5 string
	Status string
}



type Category struct {
	ID int
	Category_Name string
}


type Carts struct {
	Cartsid int
	Cartsroomid int
	User_Id int
}

type Cart_Infos struct {
	Cartsid int
	Room_Name string
	Room_Price string
	Cartsroomid int
	Cover string
	Category string
	User_Id int
	User_Name string
	

}

type Wishlists struct {
	Wishlistid int
	User_ID int
	Wishroomsid int
}

type Wishlist_Infos struct {
	Wishlistid int
	Room_Name string
	Room_Price string
	Wishroomsid int
	Cover string
	Category string
	User_Id int
	Status string
}

type Orders struct{
	Orderid int
	User_ID int
	Firstname string
	Lastname string
	Housename string
	Place string
	State string
	Mobile string
	Totalprice string
	Roomnames string
	Accountholder string
	
}

type Orderedrooms struct{
	ID int
	Roomid int
	User_Id int
	Status string
}

type Useraddress struct{
	User_Id int
	Housename string
	Place string
	Mobile string
	State string
	PIN int
}

type Addresstwo struct{
	User_Id int
	Housename string
	Place string
	Mobile string
	State string
	PIN int
}


type BookingDetails struct{
	User_Id int
	Id int
	Room_Name string
	Status string
	Cover string
	Category string
	Room_Price string
}

type ListBookings struct{
	Orderid int
	Roomid int
	User_Id int
	Totalprice string
	Status string
}

