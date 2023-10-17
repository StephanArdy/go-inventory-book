package models

type Books struct {
	ID          int64  `json:"id" form:"id" gorm:"primarykey"`        // gorm karena mau pake library gorm untuk postgre
	Title       string `json:"title" form:"title" binding:"required"` // form untuk ambil inputan dari form nanti
	Author      string `json:"author" form:"author" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	Stock       int    `json:"stock" form:"stock" binding:"required"`
}

type Login struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

const (
	USER     = "admin"
	PASSWORD = "123"
	SECRET   = "secret" // karena mau implemen JWT token
)
