package models

type Post struct {
	ID      uint   `json:"id" gorm:"primaryKey"`
	Title   string `json:"title" gorm:"not null"`
	Content string `json:"content" gorm:"not null"`
	UserID  uint   `json:"user_id" gorm:"foreignKey:ID;OnDelete:CASCADE;not null"`
	User    User   `json:"user"`
}
