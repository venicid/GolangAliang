package model

type Book struct {
	Id    int64  `gorm:"primary_key" json:"id"`
	Name  string `gorm:"not null" binding:"required" json:"name"`
	Desc  string `json:"desc"`
	Users []User `gorm:"many2many:book_users"`
}

func (receiver Book) TableName() string {
	return "book"
}
