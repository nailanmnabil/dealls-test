package entities

import "time"

type Profile struct {
	ID            string       `json:"id"`
	UserID        string    `json:"user_id"`
	Bio           string    `json:"bio"`
	Age           int       `json:"age"`
	Location      string    `json:"location"`
	ProfilePicURL string    `json:"profile_pic_url"`
	CreatedAt     time.Time `json:"created_at"`
}
