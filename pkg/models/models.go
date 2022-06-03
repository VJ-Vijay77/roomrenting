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

