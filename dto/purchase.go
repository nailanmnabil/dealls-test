package dto

type (
	PurchasePackageReq struct {
		PackageID string `json:"package_id" validate:"required,uuid"`
	}
)

type (
	GetPremiumPackagesRes struct {
		PackageID   string  `json:"package_id"`
		PackageName string  `json:"package_name"`
		Description string  `json:"description"`
		Price       float64 `json:"price"`
		FeatureType string  `json:"feature_type"`
	}
)
