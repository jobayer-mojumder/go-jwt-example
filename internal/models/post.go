package models

type Post struct {
	ID      uint   `json:"id" gorm:"primaryKey"`
	Title   string `json:"title" gorm:"required"`
	Content string `json:"content"`
	UserID  uint   `json:"user_id" gorm:"foreignKey:ID;constraint:OnUpdate:CASCADE,OnDelete:Delete"`
	User    User   `json:"user"`
}
