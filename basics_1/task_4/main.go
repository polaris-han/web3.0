package main

import (
	"main/middleware"
	"main/model"
	"main/service"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("./db/blog.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&model.Comment{}, &model.User{}, &model.Post{})

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.POST("/register", func(c *gin.Context) {
		service.Register(c, db)
	})

	router.POST("/login", func(c *gin.Context) {
		service.Login(c, db)
	})

	router.GET("/posts", func(c *gin.Context) {
		service.GetPosts(db, c)
	})

	router.GET("/posts/:id", func(c *gin.Context) {
		service.GetPostByID(db, c)
	})

	auth := router.Group("/api")
	auth.Use(middleware.JWTAuthMiddleware())
	{
		auth.POST("/posts", func(c *gin.Context) {
			service.CreatePost(db, c)
		})

		auth.PUT("/posts", func(c *gin.Context) {
			service.UpdatePost(db, c)
		})

		auth.DELETE("/posts/:id", func(c *gin.Context) {
			service.DeletePost(db, c)
		})

		auth.POST("/comments", func(c *gin.Context) {
			service.CreateComment(db, c)
		})

		auth.GET("/posts/:postID/comments", func(c *gin.Context) {
			service.GetCommentsByPostID(db, c)
		})
	}

	router.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
