package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoadBooks(e *gin.Engine) {
	e.GET("/book", BookHandler)
}

func BookHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Book Router"})
}
