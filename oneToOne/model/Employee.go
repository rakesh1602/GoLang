package model

import "gorm.io/gorm"

type Employee struct {
	gorm.Model
	EmployeeName string     `json:"employeeName"`
	EmployeeCity string     `json:"employeeCity"`
	Department   Department `json:"department"`
	DepartmentId uint       `gorm:"foreignKey:ID" json:"_"`
}
