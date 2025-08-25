package model

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title   string `gorm:"not null" json:"title"`
	Content string `gorm:"not null" json:"content"`
	UserId  uint   `gorm:"not null" json:"userId"`
	User    User   `gorm:"foreignKey:UserId"`
}

func (p *Post) InitTable(db *gorm.DB) {
	db.AutoMigrate(&Post{})
}
