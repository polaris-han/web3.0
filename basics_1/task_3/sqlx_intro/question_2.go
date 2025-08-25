package sqlx_intro

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Book struct {
	ID     int
	Title  string
	Author string
	Price  float64
}

func QueryBooksByPrice(db *sqlx.DB, price float64) ([]Book, error) {
	// 定义一个切片用于存储查询结果
	var books []Book

	// 查询语句
	query := `SELECT id, title, author, price FROM books WHERE price > ?`

	// 查询并映射到结构体切片
	err := db.Select(&books, query, price)
	if err != nil {
		return nil, fmt.Errorf("failed to query books: %w", err)
	}

	return books, nil
}
