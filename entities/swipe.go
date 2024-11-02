package entities

import "time"

const (
	RIGHT = "RIGHT"
	LEFT  = "LEFT"
)

type Swipe struct {
	ID        string    `json:"id"`
	SwiperID  string    `json:"swiper_id"`
	SwipedID  string    `json:"swiped_id"`
	SwipeType string    `json:"swipe_type"`
	SwipedAt  time.Time `json:"swiped_at"`
}
