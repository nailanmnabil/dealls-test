package entities

import "time"

const (
	UNLIMITED_SWIPE_QUOTA = "UNLIMITED_SWIPE_QUOTA"

	MONTH = "MONTH"
	YEAR  = "YEAR"
)

type PremiumPackage struct {
	ID           string    `json:"id"`
	PackageName  string    `json:"package_name"`
	Description  string    `json:"description"`
	Price        float64   `json:"price"`
	FeatureType  string    `json:"feature_type"`
	ActivePeriod string    `json:"active_period"`
	CreatedAt    time.Time `json:"created_at"`
}
