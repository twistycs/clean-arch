package models

import "time"

type Stock struct {
	ID          uint      `gorm:"primaryKey;column:ID"`
	BranchPK    uint      `gorm:"column:BRANCH_ID_PK"`
	Branch      Branch    `gorm:"foreignKey:BranchPK"`
	ProductPK   uint      `gorm:"column:PRODUCT_ID_PK"`
	Product     Branch    `gorm:"foreignKey:ProductPK"`
	Quantity    int       `gorm:"foreignKey:QUANTITY"`
	CreatedDate time.Time `gorm:"column:CREATED_DATE"`
	UpdatedDate time.Time `gorm:"column:UPDATED_DATE"`
}

// func (u *User) TableName() string {
// 	return "user"
// }
