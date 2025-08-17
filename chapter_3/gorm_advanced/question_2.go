package gorm_advanced

import (
	"gorm.io/gorm"
)

// 查询某个用户发布的所有文章及其对应的评论信息
func QueryUserPostsWithComments(db *gorm.DB, userID uint) ([]Post, error) {
	var posts []Post
	err := db.Preload("Comments").Where("user_id = ?", userID).Find(&posts).Error
	return posts, err
}

// 查询评论数量最多的文章信息
func QueryPostWithMostComments(db *gorm.DB) (Post, int64, error) {
	var post Post

	// 子查询：按评论数分组，取评论最多的文章ID
	type Result struct {
		PostID uint
		Total  int64
	}
	var result Result
	err := db.Debug().Model(&Comment{}).
		Select("post_id, COUNT(*) as total").
		Group("post_id").
		Order("total DESC").
		Limit(1).
		Scan(&result).Error
	if err != nil {
		return post, 0, err
	}

	// 查询该文章详情
	err = db.Debug().Where("id = ?", result.PostID).First(&post).Error
	return post, result.Total, err
}

// 钩子：文章创建后自动更新用户的文章数量
func (p *Post) AfterCreate(tx *gorm.DB) (err error) {
	return tx.Debug().Model(&User{}).Where("id = ?", p.UserID).
		UpdateColumn("post_count", gorm.Expr("post_count + ?", 1)).Error
}

// 钩子：评论删除后检查评论数量，若为0则更新文章评论状态
func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {
	var count int64
	tx.Debug().Model(&Comment{}).Where("post_id = ?", c.PostID).Count(&count)
	if count == 0 {
		return tx.Model(&Post{}).Where("id = ?", c.PostID).
			Update("comment_status", "无评论").Error
	}
	return nil
}
