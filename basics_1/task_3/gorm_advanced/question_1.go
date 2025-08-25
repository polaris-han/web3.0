package gorm_advanced

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name  string
	Email string
	Posts []Post // One-to-many: User has many Posts

	PostCount int // 文章数量统计字段
}

type Post struct {
	gorm.Model
	Title         string
	Content       string
	UserID        uint      // Foreign key for User
	Comments      []Comment // One-to-many: Post has many Comments
	CommentStatus string    // 评论状态字段
}

// Comment represents a comment on a post.
type Comment struct {
	gorm.Model
	Content string
	PostID  uint // Foreign key for Post
}

func AutoMigrateModels(db *gorm.DB) error {
	return db.AutoMigrate(&User{}, &Post{}, &Comment{})
}
