package model

import "time"

type CompanyBed struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	IsValid   uint      `json:"is_valid"`
	CompanyID uint      `json:"company_id"`
	BedName   string    `json:"bed_name"`
	Place     string    `json:"place"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
