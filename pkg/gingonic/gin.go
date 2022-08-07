package gingonic

import "github.com/gin-gonic/gin"

//initiating gin and returning gin engine
func GinInit()*gin.Engine{
	r := gin.Default()
	return r
}