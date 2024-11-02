package services

import (
	"context"
	"dealls-test/dto"
	"dealls-test/entities"
	"dealls-test/pkg"
	"time"

	"github.com/go-playground/validator"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Purchase interface {
	GetPremiumPackages(ctx context.Context) ([]dto.GetPremiumPackagesRes, error)
	PurchasePackage(ctx context.Context, payload dto.PurchasePackageReq, jwtPayload pkg.JwtPayload) error
}

type purchase struct {
	dbConn *gorm.DB
}

func NewPurchase(dbConn *gorm.DB) Purchase {
	return &purchase{dbConn}
}

func (p *purchase) GetPremiumPackages(ctx context.Context) ([]dto.GetPremiumPackagesRes, error) {
	packages := make([]entities.PremiumPackage, 0)
	err := p.dbConn.Model(&entities.PremiumPackage{}).Find(&packages).Error
	if err != nil {
		return nil, err
	}

	res := make([]dto.GetPremiumPackagesRes, 0)
	for _, pkg := range packages {
		res = append(res, dto.GetPremiumPackagesRes{
			PackageID:   pkg.ID,
			PackageName: pkg.PackageName,
			Description: pkg.Description,
			Price:       pkg.Price,
			FeatureType: pkg.FeatureType,
		})
	}
	return res, nil
}

func (p *purchase) PurchasePackage(ctx context.Context, payload dto.PurchasePackageReq, jwtPayload pkg.JwtPayload) error {
	validate := validator.New()
	err := validate.Struct(payload)
	if err != nil {
		return pkg.ExtendErr(pkg.ErrBadRequest, err)
	}

	var activePurchase entities.Purchase
	err = p.dbConn.Joins("JOIN premium_packages ON purchases.package_id = premium_packages.id").
		Where("purchases.user_id = ? AND premium_packages.id = ? AND purchases.expired_date > NOW()", jwtPayload.Sub, payload.PackageID).
		First(&activePurchase).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	if activePurchase.ID != "" {
		return pkg.ErrActivePackageExist
	}

	var premPackage entities.PremiumPackage
	if err := p.dbConn.Where("id = ?", payload.PackageID).First(&premPackage).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return pkg.ErrResourceNotFound
		}
		return err
	}

	monthAdder := 0
	yearAdder := 0
	if premPackage.ActivePeriod == entities.MONTH {
		monthAdder = 1
	} else if premPackage.ActivePeriod == entities.YEAR {
		yearAdder = 1
	}

	purchase := entities.Purchase{
		ID:           uuid.NewString(),
		UserID:       jwtPayload.Sub,
		PackageID:    premPackage.ID,
		PurchaseDate: time.Now().UTC(),
		ExpiredDate:  time.Now().AddDate(yearAdder, monthAdder, 0),
	}
	if err := p.dbConn.Create(&purchase).Error; err != nil {
		return err
	}

	return nil
}
