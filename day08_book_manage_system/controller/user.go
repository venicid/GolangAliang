package controller

import (
	"day08_book_manage_system/dao/mysql"
	"day08_book_manage_system/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

// 注册
func RegisterHandler(c *gin.Context) {
	p := new(model.User)
	if err := c.ShouldBindJSON(p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	// 密码进行加密：这里自己实现
	mysql.DB.Create(p)
	c.JSON(http.StatusOK, gin.H{"msg": p})
}

// 登录
func LoginHandler(c *gin.Context) {
	p := new(model.User)
	if err := c.ShouldBindJSON(p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	u := model.User{
		Username: p.Username,
		Password: p.Password,
	}
	if rows, _ := mysql.DB.Where(&u).First(&u).Rows(); rows == nil {
		c.JSON(http.StatusForbidden, gin.H{"msg": "用户名密码错误"})
		return
	}

	token := uuid.New().String()
	mysql.DB.Model(u).Update("token", token)
	c.JSON(http.StatusOK, gin.H{"token": token})
}
