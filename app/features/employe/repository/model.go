package repository

import "gorm.io/gorm"

type Employee struct {
	gorm.Model 
	FirstName       string   `json:"firstname" gorm:"type:varchar(100); not null"` 
	LastName        string   `json:"lastname" gorm:"type:varchar(100); not null"`
	HireDate        string   `json:"hiredate" gorm:"type:varchar(100); not null"`
	TerminationDate string   `json:"terminationdate" gorm:"type:varchar(100); null"`
	Salary          string   `json:"salary" gorm:"type:varchar(100); not null"`
}