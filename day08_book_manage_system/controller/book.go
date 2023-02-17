package controller

import (
	"day08_book_manage_system/dao/mysql"
	"day08_book_manage_system/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

/**
{
    "name":"认知觉醒4",
    "desc":"横空出世",
    "users": [
        {
            "id": 2
            },
             {
            "id": 3
            }
    ]
}

*/
func CreateBookHandler(c *gin.Context) {
	p := new(model.Book)

	if err := c.ShouldBindJSON(p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	mysql.DB.Create(p)
	c.JSON(http.StatusOK, gin.H{"msg": p})
}

/**
/api/v1/book
 /api/v1/book?name=被讨厌的勇气
*/
func GetBookListHandler(c *gin.Context) {

	books := []model.Book{}

	// /api/v1/book?name=被讨厌的勇气
	bookName := c.Query("name")
	if len(bookName) != 0 {
		mysql.DB.Where("name = ?", bookName).Find(&books)
		c.JSON(http.StatusOK, gin.H{"books": books})
		return
	}

	mysql.DB.Preload("Users").Find(&books)
	//mysql.DB.Find(&books) // 只查书籍，不查关联User

	c.JSON(http.StatusOK, gin.H{"books": books})
}

/**
{
            "id": 3,
            "name": "认知觉醒2",
            "desc": "改变自己的原动力",
            "Users": null
        }
*/
func UpdateBookHandler(c *gin.Context) {
	p := new(model.Book)
	if err := c.ShouldBindJSON(p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	newBook := model.Book{Id: p.Id}
	if p.Name != "" {
		newBook.Name = p.Name
	}
	if p.Desc != "" {
		newBook.Desc = p.Desc
	}

	oldBook := &model.Book{Id: p.Id}
	mysql.DB.Model(&oldBook).Updates(newBook)

	c.JSON(http.StatusOK, gin.H{"books": newBook})
}

func GetBookDetailHandler(c *gin.Context) {
	idStr := c.Param("id") // 获取URL参数

	bookId, _ := strconv.ParseInt(idStr, 10, 64)
	book := model.Book{Id: bookId}

	//mysql.DB.Preload("Users").Find(&book)
	mysql.DB.Find(&book) // 只查书籍，不查关联User

	c.JSON(http.StatusOK, gin.H{"books": book})
}

func DeleteBookHandler(c *gin.Context) {
	pipelineIdStr := c.Param("id")
	bookId, _ := strconv.ParseInt(pipelineIdStr, 10, 64)

	book := model.Book{Id: bookId}
	mysql.DB.Select("Users").Delete(&book)
	c.JSON(http.StatusOK, gin.H{"msg": "success"})
}
