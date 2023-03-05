package models

import (
	"time"
)

// User is the users model
type User struct {
	ID          int
	FirstName   string
	LastName    string
	Email       string
	Password    string
	AccessLevel int
	CreatedAt   time.Time
	UpdateAt    time.Time
}

// Room is the Room model
type Room struct {
	ID        int
	RoomName  string
	CreatedAt time.Time
	UpdateAt  time.Time
}

// Restriction is the restriction model
type Restriction struct {
	ID              int
	RestrictionName string
	CreatedAt       time.Time
	UpdateAt        time.Time
}

// Reservation is the Reservation model
type Reservation struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	RoomName  string
	Phone     string
	StartDate time.Time
	EndDate   time.Time
	RoomId    int
	Room      Room
	CreatedAt time.Time
	UpdateAt  time.Time
	Processed int
}

// RoomRestriction is the room resrtictions model
type RoomRestriction struct {
	ID            int
	ReservationID int
	RestrictionID int
	StartDate     time.Time
	EndDate       time.Time
	RoomId        int
	Room          Room
	CreatedAt     time.Time
	UpdateAt      time.Time
	Reservation   Reservation
	Restriction   Restriction
}

// MailData holds an email message
type MailData struct {
	To       string
	From     string
	Subject  string
	Content  string
	Template string
}
