package router

import "github.com/gin-gonic/gin"

func InitRouters() *gin.Engine {

	r := gin.Default()

	TestRouters(r)
	SetupApiRouters(r)

	return r
}
