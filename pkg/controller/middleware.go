package controller

import (
	"log"

	"github.com/gin-gonic/gin"
)

func UserLogedCheck(c *gin.Context) bool {

	session, err := Store.Get(c.Request, "session")
	if err != nil {
		log.Println("Couldnt find sesion!")
	}
	_, ok := session.Values["userID"]
	if !ok {
		return ok
	}
	return true
}

func AdminLogedCheck(c *gin.Context) bool {

	session, err := Store.Get(c.Request, "adminsession")
	if err != nil {
		log.Println("Couldnt find sesion!")
	}
	_, ok := session.Values["userID"]
	if !ok {
		return ok
	}
	return true
}
