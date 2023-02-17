package middleware

import (
	"day08_book_manage_system/dao/mysql"
	"day08_book_manage_system/model"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() func(c *gin.Context) {

	return func(c *gin.Context) {

		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		token := c.Request.Header.Get("token")
		if token == "" {
			c.JSON(403, gin.H{"msg": "请携带token访问"})
			c.Abort()
			return
		}

		var u model.User
		// 如果没有当前用户
		rows := mysql.DB.Where("token=?", token).First(&u).RowsAffected
		if rows != 1 {
			c.JSON(403, gin.H{"msg": "当前token错误"})
			c.Abort()
			return
		}
		// 将当前请求的userID信息保存到请求的上下文c上
		c.Set("UserId", u.Id)

		// token验证成功，返回c.Next继续，否则返回c.Abort()直接返回
		c.Next()
	}

}
