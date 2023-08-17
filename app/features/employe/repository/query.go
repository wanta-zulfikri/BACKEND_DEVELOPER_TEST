package repository

import (
	"employe/app/features/employe"
	"errors"
	"time"

	"gorm.io/gorm"
)

type EmployeRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *EmployeRepository {
	return &EmployeRepository{db: db}
}

func (er *EmployeRepository) CreateEmploye(employe employe.Core, employeID uint) error {
	var err error 
	tx := er.db.Begin()

	newEmploye := Employee {
	FirstName         : employe.FirstName,
	LastName          : employe.LastName,
	HireDate		  : employe.HireDate,
	TerminationDate   : employe.TerminationDate,
	Salary            : employe.Salary,

	}
	err = tx.Table("employe").Create(&newEmploye).Error 
	if err != nil {
		tx.Rollback()
		return err 
	}
	return err
} 

func (er *EmployeRepository) GetEmployes() ([]employe.Core, error) {
	var cores []employe.Core
	if err := er.db.Table("employe").Where("deleted_at IS NULL").Find(&cores).Error; err != nil {
		return nil, err 
	}
	return cores, nil 
} 


func (er *EmployeRepository) GetEmployeByID(employeID uint) ([]employe.Core, error) {
	var cores []employe.Core
	if err := er.db.Table("employe").Where("employe_id = ? AND deleted_at IS NULL", employeID).Find(&cores).Error; err != nil {
		return nil, err 
	}
	return cores, nil 
} 

func (er *EmployeRepository) GetEmploye(employeID uint) (employe.Core, error) {
	var input Employee
	result := er.db.Where("id = ? AND deleted_at IS NULL", employeID).Find(&input)
	if result.Error != nil {
		return employe.Core{}, result.Error	
	}
	if result.RowsAffected == 0 { 
		return employe.Core{}, result.Error
	}
		return employe.Core{
			ID: input.ID,
			FirstName:input.FirstName,
			LastName:input.LastName,
			HireDate: input.HireDate,
			TerminationDate: input.TerminationDate,
			Salary: input.Salary,

			
		}, nil
}

func (er *EmployeRepository) UpdateEmploye(id uint, updatedemploye employe.Core) error {
	if err := er.db.Model(&Employee{}).Where("id = ?", id).Updates(map[string]interface{}{
		"id"			    : updatedemploye.ID, 
		"firstname"		    : updatedemploye.FirstName,
		"lastname"			: updatedemploye.LastName,
		"hiredate"		    : updatedemploye.HireDate,
		"terminationdate"   : updatedemploye.TerminationDate, 
		"salary"            : updatedemploye.Salary,
		"updated_at"        : time.Now(),


	}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return err 
		}
		return err 
	}
	return nil 
} 

func (er *EmployeRepository) DeleteEmploye(id uint) error {
	input := Employee{}
	if err := er.db.Where("id = ?", id).Find(&input).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		return err
	}
	input.DeletedAt = gorm.DeletedAt{Time: time.Now(), Valid:true}

	if err := er.db.Save(&input).Error; err != nil {
		return err
	}
	return nil 
}