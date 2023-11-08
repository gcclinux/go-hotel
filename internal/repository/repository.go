package repository

import (
	"myapp/internal/models"
	"time"
)

type DatabaseRepo interface {
	AllUsers() bool

	InsertReservation(res models.Reservation) (int, error)
	InsertRoomRestriction(r models.RoomRestriction) error
	SearchAvailabilityByDate(start, end time.Time, roomID int) (bool, error)
}
