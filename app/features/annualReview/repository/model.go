package repository

import "gorm.io/gorm"

type AnnualReview struct {
	gorm.Model
	EmplID      string `json:"empl_id" gorm:"type:varchar(100); not null"`
	ReviewDate  string `json:"reviewdate" gorm:"type:varchar(100); not null"`
}