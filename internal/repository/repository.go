package repository

import "github.com/javy99/bookings/internal/models"

type DatabaseRepo interface {
	AllUsers() bool

	InsertReservation(res *models.Reservation) error
}
