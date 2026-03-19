package models

import "time"


type Lead struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"type:varchar(100);not null" json:"name"`
	Phone       string    `gorm:"type:varchar(20);not null" json:"phone"`
	UTMSource   string    `gorm:"type:varchar(100)" json:"utm_source"`
	UTMMedium   string    `gorm:"type:varchar(100)" json:"utm_medium"`
	UTMCampaign string    `gorm:"type:varchar(100)" json:"utm_campaign"`
	UTMTerm     string    `gorm:"type:varchar(100)" json:"utm_term"`
	UTMContent  string    `gorm:"type:varchar(100)" json:"utm_content"`
	CreatedAt   time.Time `json:"created_at"`
}