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
}



type Category struct {
	ID int
	Category_Name string
}


type Carts struct {
	ID int
	Room_Id int
	User_Id int
}
