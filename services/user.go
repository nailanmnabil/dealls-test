package services

import (
	"context"
	"dealls-test/config"
	"dealls-test/dto"
	"dealls-test/entities"
	"dealls-test/pkg"
	"time"

	"github.com/go-playground/validator"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User interface {
	Register(ctx context.Context, payload dto.RegisterReq) (dto.RegisterRes, error)
	Login(ctx context.Context, payload dto.LoginReq) (dto.LoginRes, error)
}

type user struct {
	cfg    *config.Config
	dbConn *gorm.DB
}

func NewUser(dbConn *gorm.DB, cfg *config.Config) User {
	return &user{cfg, dbConn}
}

func (u *user) Register(ctx context.Context, payload dto.RegisterReq) (dto.RegisterRes, error) {
	validate := validator.New()
	err := validate.Struct(payload)
	if err != nil {
		return dto.RegisterRes{}, pkg.ExtendErr(pkg.ErrBadRequest, err)
	}

	var existingUser entities.User
	if err := u.dbConn.Where("email = ?", payload.Email).First(&existingUser).Error; err == nil {
		return dto.RegisterRes{}, pkg.ErrEmailAlreadyRegistered
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		return dto.RegisterRes{}, err
	}

	var newUser entities.User
	err = u.dbConn.Transaction(func(tx *gorm.DB) error {
		newUser = entities.User{
			ID:       uuid.NewString(),
			Name:     payload.Name,
			Email:    payload.Email,
			Password: string(hashedPassword),
		}
		if err := tx.Create(&newUser).Error; err != nil {
			return err
		}
		if err := tx.Create(&entities.Profile{
			ID:            uuid.NewString(),
			UserID:        newUser.ID,
			Bio:           payload.Bio,
			Age:           payload.Age,
			Location:      payload.Location,
			ProfilePicURL: payload.ProfilePicURL,
		}).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return dto.RegisterRes{}, err
	}

	accessToken, _, err := pkg.GenerateToken(u.cfg.Jwt.AccessTokenSecret,
		int(time.Hour*time.Duration(u.cfg.Jwt.AccessTokenExpirationInHour)),
		pkg.JwtPayload{
			Sub: newUser.ID,
		})
	if err != nil {
		return dto.RegisterRes{}, err
	}

	return dto.RegisterRes{AccessToken: accessToken}, nil
}

func (u *user) Login(ctx context.Context, payload dto.LoginReq) (dto.LoginRes, error) {
	validate := validator.New()
	err := validate.Struct(payload)
	if err != nil {
		return dto.LoginRes{}, pkg.ExtendErr(pkg.ErrBadRequest, err)
	}

	var existingUser entities.User
	if err := u.dbConn.Where("email = ?", payload.Email).First(&existingUser).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return dto.LoginRes{}, pkg.ErrInvalidEmailOrPass
		}
		return dto.LoginRes{}, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(payload.Password)); err != nil {
		return dto.LoginRes{}, pkg.ErrInvalidEmailOrPass
	}

	accessToken, _, err := pkg.GenerateToken(u.cfg.Jwt.AccessTokenSecret,
		int(time.Hour*time.Duration(u.cfg.Jwt.AccessTokenExpirationInHour)),
		pkg.JwtPayload{
			Sub: existingUser.ID,
		})
	if err != nil {
		return dto.LoginRes{}, err
	}

	return dto.LoginRes{AccessToken: accessToken}, nil
}
