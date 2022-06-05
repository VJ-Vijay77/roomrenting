package controller

import (
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	//session,err := Store.Get(c.Request,"session")
	//if err != nil{log.Println("session not found")}
	//_,ok := session.Values["userID"]

	// if !ok {
	// 	c.Redirect(303,"/user_login")
	// 	return
	// }
	c.Redirect(303, "/user_home")
}
