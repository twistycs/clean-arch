package models

import "time"

type Product struct {
	ID           uint      `gorm:"column:ID;primaryKey"`
	ActiveStatus string    `gorm:"column:ACTIVE_STATUS"`
	Name         string    `gorm:"column:NAME"`
	ProductType  string    `gorm:"column:TYPE"`
	CreatedDate  time.Time `gorm:"column:CREATED_DATE"`
	UpdatedDate  time.Time `gorm:"column:UPDATED_DATE"`
}

// func (u *User) TableName() string {
// 	return "user"
// }
