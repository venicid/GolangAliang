package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoadUsers(e *gin.Engine) {
	e.GET("/user", UserHandler)
}

func UserHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "User Router"})
}
