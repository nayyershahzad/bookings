package dbrepo

import (
	"context"
	"time"

	"github.com/nayyershahzad/bookings/internals/models"
)

func (m *postgresDbRepo) AllUsers() bool {
	return true
}

// InsertReservation inserts the reservation into the DB
func (m *postgresDbRepo) InsertReservation(res models.Reservation) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()
	stmt := `insert into reservations (first_name, last_name, email,phone,start_date,
			end_date, room_id,start_date, end_date)
			values ($1,$2,$3,$4,$5,$6,$7,$8,$9)`
	_, err := m.DB.ExecContext(ctx, stmt,
		res.EndDate,
		res.LastName,
		res.Email,
		res.Phone,
		res.StartDate,
		res.EndDate,
		res.RoomId,
		time.Now(),
		time.Now())
	if err != nil {
		return err
	}
	return nil
}
