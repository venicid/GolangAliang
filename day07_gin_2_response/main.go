package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HelloWorldHandler(c *gin.Context) {
	c.String(http.StatusOK, "hello World!")
}

func GetBookDetailHandler(c *gin.Context) {
	bookId := c.Param("id")
	c.String(http.StatusOK, fmt.Sprintf("成功获取书籍详情：%s", bookId))
}

func GetUsersHandler(c *gin.Context) {
	username := c.Query("name")
	//username := c.DefaultQuery("name", "xxx")

	c.String(http.StatusOK, fmt.Sprintf("姓名：%s", username))
}

type Login struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

func LoginHandler(c *gin.Context) {
	var login Login

	if err := c.ShouldBind(&login); err != nil {
		// 如果数据校验不通过直接返回
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.String(http.StatusOK, fmt.Sprintf("姓名：%s, 密码：%s", login.Username, login.Password))
}

/**
响应返回
*/
func ResponseStringHandler(c *gin.Context) {
	c.String(http.StatusOK, "返回简单字符串")
}

func ResponseJsonHandler(c *gin.Context) {
	type Response struct {
		Message string      `json:"message"`
		Code    int         `json:"code"`
		Data    interface{} `json:"data"`
	}

	data := make(map[string]interface{})
	data["id"] = 1
	data["name"] = "alex"
	data["age"] = 30

	var res = Response{
		Message: "成功",
		Code:    200,
		Data:    data,
	}
	c.JSON(http.StatusOK, res)

	//// 也可以直接使用 gin.H返回 json数据
	//c.JSON(http.StatusOK, gin.H{
	// "msg": "success",
	//})
}

func ResponseRedirectHandler(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "https://www.mi.com/")

}

func main() {

	// 1.创建路由
	r := gin.Default()

	// 2.绑定路由规则，执行的函数
	/**
	路由与传参
	*/
	/**
	01 基本路由  /hello
	*/
	r.GET("/hello", HelloWorldHandler)

	/**
	02 API参数  "/book/1"
		通过context的param方法来获取
	*/
	r.GET("/book/:id", GetBookDetailHandler)

	/**
	03 URL参数  /users?name=alex
	*/
	r.GET("/users", GetUsersHandler)

	/**
	04 ShouldBind参数绑定
		以基于请求的 Content-Type 识别请求数据类型并利用反射机制
		请求中 QueryString 、 form表单 、 JSON 、 XML 等参数到结构体中

		curl --location --request POST '127.0.0.1:8080/login' \
		--header 'Content-Type: application/json' \
		--data-raw '{
			"username": "alex",
			"password":"123456"
		}'
	*/
	r.POST("/login", LoginHandler)

	/**
	响应返回
	*/
	r.GET("/response-string", ResponseStringHandler)
	r.GET("/response-json", ResponseJsonHandler)
	r.GET("/response-redirect", ResponseRedirectHandler) // 路由重定向

	// 3.监听端口，默认在8080
	fmt.Println("运行地址：http://127.0.0.1:8000")
	r.Run(":8080")
}
