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

type Profile interface {
	View(ctx context.Context, jwtPayload pkg.JwtPayload) (dto.ViewRes, error)
	Swipe(ctx context.Context, payload dto.SwipeReq, jwtPayload pkg.JwtPayload) error
}

type profile struct {
	dbConn *gorm.DB
}

func NewProfile(dbConn *gorm.DB) Profile {
	return &profile{dbConn}
}

func (p *profile) View(ctx context.Context, jwtPayload pkg.JwtPayload) (dto.ViewRes, error) {
	conn := p.dbConn
	today := time.Now().UTC().Format("2006-01-02")

	// Check if the user has already viewed a profile but hasn't swiped it yet
	var lastProfileVisit entities.ProfileVisit
	err := conn.Where("visitor_id = ? AND visit_date = ? AND swiped_at IS NULL", jwtPayload.Sub, today).Last(&lastProfileVisit).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return dto.ViewRes{}, err
	}
	if lastProfileVisit.ID != "" {
		var lastViewedUser entities.User
		err = conn.Preload("Profile").Where("id = ?", lastProfileVisit.VisitedID).First(&lastViewedUser).Error
		if err != nil {
			return dto.ViewRes{}, err
		}

		return dto.ViewRes{
			UserID:         lastViewedUser.ID,
			ProfileVisitID: lastProfileVisit.ID,
			Name:           lastViewedUser.Name,
			Bio:            lastViewedUser.Profile.Bio,
			Age:            lastViewedUser.Profile.Age,
			Location:       lastViewedUser.Profile.Location,
			ProfilePicURL:  lastViewedUser.Profile.ProfilePicURL,
		}, nil
	}

	// Check if the user have active UNLIMITED_SWIPE_QUOTA premium package
	hasPremiumPackage := true

	var purchase entities.Purchase
	err = conn.
		Joins("JOIN premium_packages ON purchases.package_id = premium_packages.id").
		Where("purchases.user_id = ? AND premium_packages.feature_type = ? AND purchases.expired_date > NOW()", jwtPayload.Sub, entities.UNLIMITED_SWIPE_QUOTA).
		First(&purchase).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return dto.ViewRes{}, err
	}

	if err == gorm.ErrRecordNotFound {
		hasPremiumPackage = false
	}

	// Check the user swipe quota
	if !hasPremiumPackage {
		var visitCount int64
		if err := conn.Model(&entities.Swipe{}).Where("swiper_id = ? AND DATE(swiped_at) = CURRENT_DATE", jwtPayload.Sub).
			Count(&visitCount).Error; err != nil {
			return dto.ViewRes{}, err
		}

		if visitCount >= 10 {
			return dto.ViewRes{}, pkg.ErrLimitReached
		}
	}

	// Get random profile
	var randomUser entities.User
	err = conn.Model(&entities.User{}).Preload("Profile").
		Joins("LEFT JOIN profile_visits ON profile_visits.visited_id = users.id AND profile_visits.visitor_id = ?", jwtPayload.Sub).
		Where("users.id != ?", jwtPayload.Sub).
		Where("profile_visits.visited_id IS NULL").
		Order("RANDOM()").
		Limit(1).
		Find(&randomUser).Error
	if err != nil {
		return dto.ViewRes{}, err
	}
	if randomUser.ID == "" {
		return dto.ViewRes{}, pkg.ErrNoUserLeft
	}

	// Save visit data
	profileVisit := entities.ProfileVisit{
		ID:        uuid.NewString(),
		VisitorID: jwtPayload.Sub,
		VisitedID: randomUser.ID,
		VisitDate: time.Now().UTC(),
	}
	err = conn.Create(&profileVisit).Error
	if err != nil {
		return dto.ViewRes{}, err
	}

	return dto.ViewRes{
		UserID:         randomUser.ID,
		ProfileVisitID: profileVisit.ID,
		Name:           randomUser.Name,
		Bio:            randomUser.Profile.Bio,
		Age:            randomUser.Profile.Age,
		Location:       randomUser.Profile.Location,
		ProfilePicURL:  randomUser.Profile.ProfilePicURL,
	}, nil
}

func (p *profile) Swipe(ctx context.Context, payload dto.SwipeReq, jwtPayload pkg.JwtPayload) error {
	validate := validator.New()
	err := validate.Struct(payload)
	if err != nil {
		return pkg.ExtendErr(pkg.ErrBadRequest, err)
	}

	return p.dbConn.Transaction(func(tx *gorm.DB) error {
		var profileVisit entities.ProfileVisit
		err = tx.Where("id = ?", payload.ProfileVisitID).First(&profileVisit).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return pkg.ErrResourceNotFound
			}
			return err
		}

		if profileVisit.SwipedAt != nil {
			return pkg.ErrProfileAlreadySwiped
		}

		now := time.Now().UTC()
		swipe := entities.Swipe{
			ID:        uuid.NewString(),
			SwiperID:  jwtPayload.Sub,
			SwipedID:  profileVisit.VisitedID,
			SwipeType: payload.SwipeType,
			SwipedAt:  now,
		}
		err = p.dbConn.Create(&swipe).Error
		if err != nil {
			return err
		}

		return tx.Model(&entities.ProfileVisit{}).Where("id = ?", payload.ProfileVisitID).Update("swiped_at", now).Error
	})
}
