package employe

import "github.com/labstack/echo/v4"

type Core struct {
	ID        uint      
	FirstName string 
	LastName  string 
	HireDate  string 
	TerminationDate string
	Salary    string 
}  

type Repository interface {
	CreateEmploye(employe Core, employeID uint) error
	GetEmployes() ([]Core, error)
	GetEmployeByID(employeid uint) ([]Core, error)
	GetEmploye(employeID uint) (Core, error)
	UpdateEmploye(id uint, updatedEmploye Core) error
	DeleteEmploye(id uint) error
}

type Service interface {
	CreateEmploye(employe Core, employeID uint) error
	GetEmployes() ([]Core, error)
	GetEmployeByID(employeid uint) ([]Core, error)
	GetEmploye(employeID uint) (Core, error)
	UpdateEmploye(id uint, updatedEmploye Core) error
	DeleteEmploye(id uint) error
}

type Handler interface {
	CreateEmploye() echo.HandlerFunc 
	GetEmploye() echo.HandlerFunc
	GetEmployeByID() echo.HandlerFunc
	GetEmployes() echo.HandlerFunc
	UpdateEmploye() echo.HandlerFunc
	DeleteEmploye() echo.HandlerFunc
}