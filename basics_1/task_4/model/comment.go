package model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Content string `gorm:"not null" json:"content"`
	PostId  uint   `json:"postId"`
	Post    Post   `gorm:"foreignKey:PostId"`
	UserId  uint   `json:"userId"`
	User    User   `gorm:"foreignKey:UserId"`
}
