package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
Engine 容器对象,整个框架的基础
Engine.trees 负责存储路由和handle方法的映射,采用类似字典树的结构
Engine.RouterGroup ,其中的Handlers存储着所有中间件
Context 上下文对象,负责处理 请求和回应 ,其中的 handlers 是存储处理请求时中间件和处理方法的
*/

func main() {

	// 1.创建 (实例化gin.Engine结构体对象）
	r := gin.Default()

	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello World!")
	})

	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8080")
	r.Run(":8000")
}
