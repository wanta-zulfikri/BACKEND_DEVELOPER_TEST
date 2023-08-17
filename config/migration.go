package config 

import (
	employe "employe/app/features/employe/repository"

	"gorm.io/gorm"
)
func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&employe.Employee{})
	return err
}