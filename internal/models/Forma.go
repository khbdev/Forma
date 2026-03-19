package models

import "time"


type Lead struct {
	ID          uint      `gorm:"primaryKey"`
	Name        string    `gorm:"type:varchar(100);not null"`
	Phone       string    `gorm:"type:varchar(20);not null"`
	UTMSource   string    `gorm:"type:varchar(100)"`
	UTMMedium   string    `gorm:"type:varchar(100)"`
	UTMCampaign string    `gorm:"type:varchar(100)"`
	UTMTerm     string    `gorm:"type:varchar(100)"`
	UTMContent  string    `gorm:"type:varchar(100)"`
	CreatedAt   time.Time
}