package model

type User struct {
	Id       int64  `gorm:"primary_key" json:"id"`
	Username string `gorm:"not null" binding:"required" json:"username"`
	Password string `gorm:"not null" binding:"required" json:"password"`
	Token    string `json:"token"`
}

func (receiver User) TableName() string {
	return "user"
}
