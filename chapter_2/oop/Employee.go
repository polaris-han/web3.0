package oop

import "fmt"

type Employee struct {
	Person
	EmployeeID string
}

func (e Employee) PrintInfo() {
	fmt.Println("age:", e.Age, "\nname:", e.Name, "\nemployeeID:", e.EmployeeID)
}
