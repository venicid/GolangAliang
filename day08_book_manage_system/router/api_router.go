package router

import (
	"day08_book_manage_system/controller"
	"day08_book_manage_system/middleware"
	"github.com/gin-gonic/gin"
)

func SetupApiRouters(r *gin.Engine) {
	r.POST("/register", controller.RegisterHandler)
	r.POST("/login", controller.LoginHandler)

	v1 := r.Group("/api/v1")
	v1.Use(middleware.AuthMiddleware()) // 中间件验证
	v1.POST("book", controller.CreateBookHandler)
	v1.GET("book", controller.GetBookListHandler)
	v1.GET("book/:id", controller.GetBookDetailHandler)
	v1.PUT("book", controller.UpdateBookHandler)
	v1.DELETE("book/:id", controller.DeleteBookHandler)
}
