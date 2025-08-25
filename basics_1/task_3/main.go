package main

import (
	"chapter_3/gorm_advanced"
	"chapter_3/sql_practice"
	"chapter_3/sqlx_intro"
	"fmt"
	"math/rand"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func sqlPractice() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	stu := sql_practice.Student{}
	stu.InitTable(db)
	stu.Insert(db)
	stu.UpdateStu(db)
	stu.SelectStu(db)
	stu.DeleteStu(db)

	sql_practice.InsertRandomAccounts(db, 10)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	fromAccountID := uint(r.Intn(10) + 1) // 随机选择一个账户ID
	fmt.Println("转出账户ID:", fromAccountID)
	toAccountID := uint(r.Intn(10) + 1) // 随机选择一个账户ID
	fmt.Println("转入账户ID:", toAccountID)
	amount := r.Float64() * 1000 // 随机选择一个转账金额
	fmt.Println("转账金额:", amount)

	sql_practice.Transfer(db, fromAccountID, toAccountID, amount)
}

func sqlxIntro() {
	db, err := sqlx.Connect("sqlite3", "test.db")
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	result1, err := sqlx_intro.QueryByDept("技术部", db)
	fmt.Println(result1)

	if err != nil {
		panic("failed to query database")
	}

	result2, err := sqlx_intro.QueryHighestPaidEmployee(db)
	if err != nil {
		panic("failed to query database")
	}
	fmt.Println(result2)

	books, err := sqlx_intro.QueryBooksByPrice(db, 50.0)
	if err != nil {
		panic("failed to query database")
	}
	fmt.Println(books)
}

func gormAdvanced() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	gorm_advanced.AutoMigrateModels(db)

	// posts1, total, err := gorm_advanced.QueryPostWithMostComments(db)
	// fmt.Println("评论最多的文章:", posts1, "评论数量:", total)

	// gorm_advanced.QueryUserPostsWithComments(db, 1)

	// user := gorm_advanced.User{Name: "张三", Email: "zhangsan@example.com"}
	// db.Create(&user)

	// post := gorm_advanced.Post{Title: "第一篇文章", Content: "这是我的第一篇文章", UserID: user.ID}
	// db.Create(&post)

	// comment := gorm_advanced.Comment{Content: "这是一条评论", PostID: post.ID}
	// db.Create(&comment)

	comment := gorm_advanced.Comment{Model: gorm.Model{ID: 8}}
	db.Debug().Model(&comment).Find(&comment)
	db.Debug().Model(&comment).Delete(&comment)
}

func main() {
	// sqlPractice()

	// sqlxIntro()

	gormAdvanced()

}
