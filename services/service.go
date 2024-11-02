package services

import (
	"dealls-test/config"

	"gorm.io/gorm"
)

type Service struct {
	User     User
	Profile  Profile
	Purchase Purchase
}

func NewService(dbConn *gorm.DB, cfg *config.Config) *Service {
	userSvc := NewUser(dbConn, cfg)
	profileSvc := NewProfile(dbConn)
	purchaseSvc := NewPurchase(dbConn)

	return &Service{
		User:     userSvc,
		Profile:  profileSvc,
		Purchase: purchaseSvc,
	}
}
