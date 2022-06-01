package gingonic

import "github.com/gin-gonic/gin"

func GinInit()*gin.Engine{
	r := gin.Default()
	return r
}