package model

import "gorm.io/gorm"

type User struct {
	Username string `gorm:"unique;not null" json:"username"`
	Email    string `gorm:"unique;not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
	gorm.Model
}

func (u *User) InitTable(db *gorm.DB) {
	db.AutoMigrate(&User{})
}
