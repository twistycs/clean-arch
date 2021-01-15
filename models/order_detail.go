package models

import "time"

type OrderDetail struct {
	ID          uint      `gorm:"column:ID;primaryKey"`
	OrderNoFK   int       `gorm:"column:ORDER_NO_FK"`
	ProductPK   uint      `gorm:"column:PRODUCT_ID_FK;" json:"-"`
	Product     Product   `gorm:"foreignKey:ProductPK"`
	CreatedDate time.Time `gorm:"column:CREATED_DATE"`
	UpdatedDate time.Time `gorm:"column:UPDATED_DATE"`
}

// func (o *Order) TableName() string {
// 	return "order"
// }
