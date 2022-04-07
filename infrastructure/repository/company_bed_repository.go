package repository

import (
	"github.com/goki0524/clean-arch/domain/model"
	"github.com/jinzhu/gorm"
)

type ICompanyBedRepository interface {
	FindOne(uint) *model.CompanyBed
}

type CompanyBedRepository struct {
	db *gorm.DB
}

func NewCompanyBedRepository(db *gorm.DB) *CompanyBedRepository {
	return &CompanyBedRepository{db: db}
}

const tableName = "company_beds"

func (repo *CompanyBedRepository) FindOne(ID uint) *model.CompanyBed {
	companyBed := model.CompanyBed{}
	if repo.db.Table(tableName).Where(model.CompanyBed{ID: ID}).First(&companyBed).RecordNotFound() {
		return nil
	}

	return &companyBed
}
