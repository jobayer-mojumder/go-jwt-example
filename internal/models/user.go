package models

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Name     string `json:"name" gorm:"not null"`
	Email    string `json:"email" gorm:"not null;unique"`
	Password []byte `json:"password" gorm:"not null"`
}
