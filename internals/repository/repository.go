package repository

import "github.com/nayyershahzad/bookings/internals/models"

type DatabaseRepo interface {
	AllUsers() bool
	InsertReservation(res models.Reservation) error
}
