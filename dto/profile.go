package dto

type (
	SwipeReq struct {
		SwipeType      string `json:"swipe_type" validate:"required,oneof=LEFT RIGHT"`
		ProfileVisitID string `json:"profile_visit_id" validate:"required,uuid"`
	}
)

type (
	ViewRes struct {
		UserID         string `json:"user_id"`
		ProfileVisitID string `json:"profile_visit_id"`
		Name           string `json:"name"`
		Bio            string `json:"bio"`
		Age            int    `json:"age"`
		Location       string `json:"location"`
		ProfilePicURL  string `json:"profile_pic_url"`
	}
)
