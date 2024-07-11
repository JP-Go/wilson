package entities

import (
	vo "github.com/JP-Go/wilson/backend/internal/domain/value_objects"
	"github.com/google/uuid"
)

type ParticipantId = uuid.NullUUID

type Participant struct {
	Id                ParticipantId `json:"id"`
	Name              string        `json:"name"`
	Email             vo.Email      `json:"email"`
	ConfirmedPresence bool          `json:"confirmed_presence"`
}

type NewParticipantParams struct {
	id                uuid.NullUUID
	name              string
	email             vo.Email
	confirmedPresence bool
}

func (p *Participant) IsEqual(pother Participant) bool {
	return p.Email.Value == pother.Email.Value
}

func (p *Participant) ConfirmPresence() {
	p.ConfirmedPresence = true
}

func (p *Participant) ChangeEmail(email vo.Email) {
	p.Email = email
}

func NewParticipant(params NewParticipantParams) (*Participant, error) {
	err := params.email.ValidateEmail()
	if err != nil {
		return nil, err
	}
	participant := Participant{
		Id: uuid.NullUUID{
			UUID:  uuid.New(),
			Valid: true,
		},
		Name:              params.name,
		Email:             params.email,
		ConfirmedPresence: params.confirmedPresence,
	}
	if params.id.Valid {
		participant.Id = params.id
	}
	return &participant, nil
}
