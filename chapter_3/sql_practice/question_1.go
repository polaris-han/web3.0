package sql_practice

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name  string
	Age   int
	Grade string
}

func (s *Student) InitTable(db *gorm.DB) {
	db.Debug().AutoMigrate(&Student{})
}

// 编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"
func (s *Student) Insert(db *gorm.DB) error {
	if db == nil {
		return errors.New("数据库连接不能为空")
	}
	var stu Student = Student{Name: "张三", Age: 20, Grade: "三年级"}

	db.Debug().Create(&stu)

	return nil
}

// 编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息
func (s *Student) SelectStu(db *gorm.DB) error {
	if db == nil {
		return errors.New("数据库连接不能为空")
	}

	var results []Student = []Student{}
	db.Debug().Where("age > ?", 18).Find(&results)
	fmt.Println("students 表中所有年龄大于 18 岁的学生信息", results)

	return nil
}

// 编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"
func (s *Student) UpdateStu(db *gorm.DB) error {
	if db == nil {
		return errors.New("数据库连接不能为空")
	}

	db.Debug().Model(&Student{}).Where("name = ?", "张三").Update("Grade", "四年级")

	return nil
}

// 编写SQL语句删除 students 表中年龄小于 15 岁的学生记录
func (s *Student) DeleteStu(db *gorm.DB) error {
	if db == nil {
		return errors.New("数据库连接不能为空")
	}

	db.Debug().Where("age < ?", "15").Delete(&Student{})

	return nil
}
