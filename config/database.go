package config 

import (
	"fmt"
	"time"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)



func GetConnection(config Configuration) (*gorm.DB, error) {
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Database.Username,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.Name,
	)

	db, err := gorm.Open(mysql.Open(dataSource), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}
	sqlDB.SetMaxIdleConns(1)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxIdleTime(1 * time.Minute)
	sqlDB.SetConnMaxIdleTime(1 * time.Minute) 
	err = sqlDB.Ping() 
	if err != nil {
		return nil, fmt.Errorf("unable to establish a good connection to the database: %v", err)
	}
	return db, nil
}
