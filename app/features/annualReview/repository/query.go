package repository

import (
	annualreview "employe/app/features/annualReview"
	"errors"
	"time"

	"gorm.io/gorm"
)

type AnnualReviewRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *AnnualReviewRepository {
	return &AnnualReviewRepository{db: db}
}

func (ar *AnnualReviewRepository) CreateAnnual(annualreview annualreview.Core, EmplID uint) error {
	var err error 
	tx := ar.db.Begin()

	newAnnualReview := AnnualReview {
		EmplID : annualreview.EmplID,
		ReviewDate : annualreview.ReviewDate,
	} 

	err = tx.Table("annualreview").Create(&newAnnualReview).Error 
	if err != nil {
		tx.Rollback()
		return err
	} 
	return err 
} 

func (ar *AnnualReviewRepository) GetAnnuals() ([]annualreview.Core, error) {
	var cores []annualreview.Core 
	if err := ar.db.Table("annualreview").Where("deleted_at IS NULL").Find(&cores).Error; err != nil {
		return nil, err 
	}
	return cores, nil 
}

func (ar *AnnualReviewRepository) GetEmplByID(EmplID uint) ([]annualreview.Core, error) {
	var cores []annualreview.Core 
	if err := ar.db.Table("annualreview").Where("empl_id = ? AND deleted_at IS NULL", EmplID).Find(&cores).Error; err != nil {
		return nil, err
	}
	return cores, nil 
}

func (ar *AnnualReviewRepository)GetAnnual(EmplID uint ) (annualreview.Core, error) {
	var input AnnualReview 
	result := ar.db.Where("id = ? AND deleted_at IS NULL", EmplID).Find(&input) 
	if result.Error !=  nil {
		return annualreview.Core{}, result.Error
	}
	if result.RowsAffected == 0 {
		return annualreview.Core{}, result.Error
	}

		return annualreview.Core{
			EmplID 	  	  : input.EmplID ,
			ReviewDate    : input.ReviewDate,
        }, nil
} 

func (ar *AnnualReviewRepository) UpdateAnnual(id uint, updatedEmplID annualreview.Core) error {
	if err := ar.db.Model(&AnnualReview{}).Where("id = ?" , id).Updates(map[string]interface{}{
		"emplid" : updatedEmplID.EmplID, 
		"reviewdate":updatedEmplID.ReviewDate,
		"updated_at" : time.Now(),
		}).Error ; err!=nil { 
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return err 
			}
			return err
		}
		return nil
} 

func (ar *AnnualReviewRepository) DeleteAnnual(id uint) error {
	input := AnnualReview{} 
	if err := ar.db.Where("id = ?", id).Find(&input).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return err 
		}
		return err 
	}
	input.DeletedAt = gorm.DeletedAt{Time: time.Now(), Valid: true} 

	if err:= ar.db.Save(&input).Error; err!= nil {
		return err
	} 
	return nil
}
