package entities

import "time"

type Purchase struct {
	ID             string         `json:"id"`
	UserID         string         `json:"user_id"`
	PackageID      string         `json:"package_id"`
	PurchaseDate   time.Time      `json:"purchase_date"`
	ExpiredDate    time.Time      `json:"expired_date"`
	PremiumPackage PremiumPackage `json:"premium_package" gorm:"foreignKey:PackageID;references:ID"`
}
