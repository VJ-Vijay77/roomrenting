package controllertwo

import "github.com/gin-gonic/gin"


func Wallet ( c *gin.Context){

	c.HTML(200,"wallet.gohtml",gin.H{
		
	})
}