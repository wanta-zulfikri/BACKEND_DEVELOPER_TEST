package annualreview

import "github.com/labstack/echo/v4"

type Core struct {
	ID          uint 
	EmplID      string 
	ReviewDate  string 
}

type Repository interface {
	CreateAnnual(annualreview Core, EmplID uint) error
	GetAnnual(EmplID uint) (Core, error)
	GetEmplByID(EmplID uint) ([]Core, error)
	GetAnnuals([]Core, error)
	UpdateAnnual(id uint, updatedEmplID Core) error
	DeleteAnnual(id uint) error
}

type Service interface {
	CreateAnnual(annualreview Core, EmplID uint) error
	GetAnnual(EmplID uint) (Core, error)
	GetEmplByID(EmplID uint) ([]Core, error)
	GetAnnuals([]Core, error) 
	UpdateAnnual(id uint, updatedEmplID Core) error
	DeleteAnnual(id uint) error
}

type Handler interface {
	CreateAnnual() echo.HandlerFunc 
	GetAnnual()    echo.HandlerFunc
	GetEmplByID()  echo.HandlerFunc 
	GetAnnuals()   echo.HandlerFunc 
	UpdateAnnual() echo.HandlerFunc
	DeleteAnnual() echo.HandlerFunc 
}