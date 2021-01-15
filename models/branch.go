package models

import "time"

type Branch struct {
	ID           uint      `gorm:"primaryKey;column:ID"`
	ActiveStatus string    `gorm:"column:ACTIVE_STATUS;size:1"`
	Name         string    `gorm:"column:NAME"`
	Address      string    `gorm:"column:ADDRESS"`
	Province     string    `gorm:"column:PROVINCE"`
	ZipCode      string    `gorm:"column:ZIP_CODE"`
	CreatedDate  time.Time `gorm:"column:CREATED_DATE"`
	UpdatedDate  time.Time `gorm:"column:UPDATED_DATE"`
}

// func (u *User) TableName() string {
// 	return "user"
// }
