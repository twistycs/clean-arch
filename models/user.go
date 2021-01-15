package models

type User struct {
	ID               uint   `gorm:"primaryKey;column:ID"`
	UserID           string `gorm:"column:USER_ID"`
	FirstName        string `gorm:"column:FIRST_NAME"`
	LastName         string `gorm:"column:LAST_NAME"`
	Email            string `gorm:"column:EMAIL"`
	MobileNo         string `gorm:"column:MOBILE_NO"`
	PresentAddress   string `gorm:"column:PRESENT_ADDRESS"`
	PresentProvince  string `gorm:"column:PRESENT_PROVINCE"`
	PresentZipCode   string `gorm:"column:PRESENT_ZIP_CODE"`
	DeliveryAddress  string `gorm:"column:DELIVERY_ADDRESS"`
	DeliveryProvince string `gorm:"column:DELIVERY_PROVINCE"`
	DeliveryZipCode  string `gorm:"column:DELIVERY_ZIP_CODE"`
}

// func (u *User) TableName() string {
// 	return "user"
// }
