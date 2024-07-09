package entities

import (
	"errors"
	"time"

	valueobjects "github.com/JP-Go/wilson/backend/internal/domain/value_objects"
	"github.com/google/uuid"
)

type TripId = uuid.NullUUID

type Trip struct {
	Id           TripId             `json:"id"`
	StartsAt     time.Time          `json:"starts_at"`
	EndsAt       time.Time          `json:"ends_at"`
	Destination  string             `json:"destination"`
	OwnerEmail   valueobjects.Email `json:"owner_email"`
	OwnerName    string             `json:"owner_name"`
	Participants []Participant      `json:"participants"`
}

type TripModel struct {
	Id          uuid.UUID
	StartsAt    time.Time
	EndsAt      time.Time
	Destination string
	OwnerEmail  string
	OwnerName   string
}

type TripConfig struct {
	startsAt    time.Time
	endsAt      time.Time
	destination string
	ownerName   string
	ownerEmail  string
}

func NewTrip(id uuid.NullUUID, config TripConfig) (*Trip, error) {
	if config.endsAt.Before(config.startsAt) {
		return nil, errors.New("End date should be after the start date")
	}
	onwerEmail, err := valueobjects.NewEmail(config.ownerEmail)
	if err != nil {
		return nil, err
	}

	trip := Trip{
		Id: uuid.NullUUID{
			UUID:  uuid.New(),
			Valid: true,
		},
		StartsAt:     config.startsAt,
		EndsAt:       config.endsAt,
		Destination:  config.destination,
		OwnerEmail:   *onwerEmail,
		OwnerName:    config.ownerName,
		Participants: make([]Participant, 5),
	}
	if id.Valid {
		trip.Id = id
	}
	return &trip, nil
}

func (t *Trip) Invite(p Participant) error {
	for _, participant := range t.Participants {
		if participant.IsEqual(p) {
			return errors.New("Person already invited")
		}
	}
	t.Participants = append(t.Participants, p)
	return nil
}
