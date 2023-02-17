package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HelloHandler(c *gin.Context) {
	fmt.Println("执行了Get方法")
	c.String(http.StatusOK, "hello World!")
}

func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("我是一个全局中间件")
	}
}

func MiddleWareNext() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("开始执行中间件")
		c.Next()
		fmt.Println("视图函数执行完成后再返回到next()方法")
	}

}

func LoginHandler(c *gin.Context) {
	c.String(http.StatusOK, "login页面，请登录")
}

func HomeHandler(c *gin.Context) {
	c.String(http.StatusOK, "home页面，欢迎用户")
}

//func AuthMiddleware() gin.HandlerFunc {
func AuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// token验证成功，返回c.Next继续，
		//  否则返回c.Abort()直接返回

		token := c.Request.Header.Get("token")
		fmt.Println("token: ", token) // 获取token： xxxxxx
		if token == "" {
			c.String(http.StatusUnauthorized, "身份验证不通过")
			c.Abort()
			return
		}
		c.Next()
	}
}

/**
中间件
	在处理请求的过程中，加入用户自己的钩子（Hook）函数,这个钩子函数就叫中间件
	中间件适合处理一些公共的业务逻辑比如 登录认证、权限校验 、数据分页、记录日志、耗时统计等
*/
func main() {

	r := gin.Default()

	/**
	全局中间件
	*/
	//r.Use(MiddleWare())
	//r.GET("/hello", HelloHandler)
	// 我是一个全局中间件
	//执行了Get方法

	/**
	局部中间件
	*/
	r.GET("/hello", MiddleWare(), HelloHandler)

	/**
	Next方法
		调用next() 方法，会从next()方法调用的地方跳转到视图函数，视图函数执行完成
		再返回到next() 方法处
	*/
	r.GET("/hello2", MiddleWareNext(), HelloHandler)
	// 开始执行中间件
	//执行了Get方法
	//视图函数执行完成后再返回到next()方法

	/**
	token认证
	*/
	// 登录页面无需验证
	r.GET("/login", LoginHandler)
	// home页需要用户登录才能查看

	r.GET("/home", AuthMiddleware(), HomeHandler)

	r.Run(":8080")

}
