package models

type Post struct {
	ID      uint   `json:"id" gorm:"primaryKey"`
	Title   string `json:"title"`
	Content string `json:"content"`
	UserID  uint   `json:"user_id" gorm:"foreignKey:ID" binding:"required" form:"user_id" query:"user_id" uri:"user_id"`
}
