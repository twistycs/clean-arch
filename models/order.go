package models

import (
	"time"
)

type Order struct {
	ID           uint          `gorm:"column:ORDER_ID;primaryKey"`
	OrderNo      uint          `gorm:"column:ORDER_NO;primaryKey;index"`
	OrderDetails []OrderDetail `gorm:"foreignKey:OrderNoFK;references:OrderNo"`
	UserOrderPK  uint          `gorm:"column:USER_ID_PK;" json:"-"`
	User         User          `gorm:"foreignKey:UserOrderPK"`
	BranchPK     uint          `gorm:"column:BRANCH_ID_PK" json:"-"`
	Branch       Branch        `gorm:"foreignKey:BranchPK"`
	TotalAmount  float64       `gorm:"column:TOTAL_AMOUNT"`
	Remark       float64       `gorm:"column:REMARK"`
	CreatedDate  time.Time     `gorm:"column:CREATED_DATE"`
	UpdatedDate  time.Time     `gorm:"column:UPDATED_DATE"`
}

// func (o *Order) TableName() string {
// 	return "order"
// }
