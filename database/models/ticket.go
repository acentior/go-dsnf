package models

import (
	"time"
)

type Ticket struct {
	ID                       uint      `json:"id"`
	UserId                   *User     `json:"user_id"`
	Salestime                time.Time `json:"sales_time"`
	DateIssued               string    `json:"date_issued"`
	ScheduleDepartureStation *Station  `json:"schedule_departure_station"`
	ScheduleArrivalStation   *Station  `json:"schedule_arrival_station"`
	IsFirstClass             bool      `json:"is_first_class"`
	IsSecondClass            bool      `json:"is_second_class"`
	SeatReserved             string    `json:"seat_reserved"`
	Price                    float64   `json:"price"`
}
