package services

import (
	"employe/app/features/employe"
	"errors"

	"gorm.io/gorm"
)

type EmployeService struct {
	r employe.Repository
}

func New(r employe.Repository) employe.Service {
	return &EmployeService{r: r}
}

func (s *EmployeService) CreateEmploye(employe employe.Core, employeID uint) error {
	err := s.r.CreateEmploye(employe, employeID)
	if err != nil {
		return err
	}
	return nil 
} 

func (es *EmployeService) GetEmployes() ([]employe.Core, error) {
	employe, err := es.r.GetEmployes()
	if err != nil {
		return nil, err
	}
	return employe, nil
} 

func (es *EmployeService) GetEmployeByID(employeID uint) ([]employe.Core, error) {
	employe, err := es.r.GetEmployeByID(employeID)
	if err != nil {
		return nil, err 
	}
	return employe, nil
} 

func (es *EmployeService) GetEmploye(employeID uint) (employe.Core, error) {
	employes, err := es.r.GetEmploye(employeID)
	if err != nil {
		return employe.Core{}, err
	}
	return employes, nil
}

func (es *EmployeService) UpdateEmploye(id uint, updatedEmploye employe.Core) error {
	updatedEmploye.ID = id
	if err := es.r.UpdateEmploye(id, updatedEmploye); err != nil {
		return nil 
	}
	return nil
}

func (es *EmployeService) DeleteEmploye(id uint) error {
	err := es.r.DeleteEmploye(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return err 
		}
		return err 
	}
	return nil
}