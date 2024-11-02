package dto

type (
	RegisterReq struct {
		Name          string `json:"name" validate:"required,min=3"`
		Email         string `json:"email" validate:"required,email"`
		Password      string `json:"password" validate:"required,min=8,max=16"`
		Bio           string `json:"bio" validate:"required,min=15"`
		Age           int    `json:"age" validate:"required,min=0"`
		Location      string `json:"location" validate:"required"`
		ProfilePicURL string `json:"profile_pic_url" validate:"required,url"`
	}
	LoginReq struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=8,max=16"`
	}
)

type (
	RegisterRes struct {
		AccessToken string `json:"access_token"`
	}
	LoginRes struct {
		AccessToken string `json:"access_token"`
	}
)
