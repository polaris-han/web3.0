package sqlx_intro

import (
	"github.com/jmoiron/sqlx"
)

type Employee struct {
	ID         int
	Name       string
	Department string
	Salary     float64
}

func QueryByDept(department string, db *sqlx.DB) ([]Employee, error) {
	var employees []Employee
	err := db.Select(&employees, "SELECT * FROM employees WHERE department = ?", department)
	return employees, err
}

func QueryHighestPaidEmployee(db *sqlx.DB) (Employee, error) {
	var employee Employee
	err := db.Get(&employee, "SELECT * FROM employees ORDER BY salary DESC LIMIT 1")
	return employee, err
}
