package entities

import "time"

type ProfileVisit struct {
	ID        string     `json:"id"`
	VisitorID string     `json:"visitor_id"`
	VisitedID string     `json:"visited_id"`
	VisitDate time.Time  `json:"visit_date"`
	SwipedAt  *time.Time `json:"swiped_at"`
}
