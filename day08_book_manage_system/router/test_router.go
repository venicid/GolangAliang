package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func TestHandler(c *gin.Context) {
	c.String(http.StatusOK, "ok")

}

// 测试路由访问： http://127.0.0.1:8888/api/v1/test
func TestRouters(r *gin.Engine) {
	v1 := r.Group("/api/v1")

	v1.GET("test", TestHandler)
}
