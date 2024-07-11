package mappers

import (
	"github.com/JP-Go/wilson/backend/infra/database/pgstore"
	"github.com/JP-Go/wilson/backend/internal/domain/entities"
	vo "github.com/JP-Go/wilson/backend/internal/domain/value_objects"
	"github.com/google/uuid"
)

func ParticipantFromDBToDomain(dbParticipant pgstore.Participant) *entities.Participant {
	return &entities.Participant{
		Id: uuid.NullUUID{
			UUID:  dbParticipant.ID,
			Valid: true,
		},
		Email: vo.Email{
			Value: dbParticipant.Email,
			Valid: true,
		},
		ConfirmedPresence: dbParticipant.IsConfirmed,
		TripId: uuid.NullUUID{
			UUID:  dbParticipant.TripID,
			Valid: true,
		},
	}
}

func ParticipantFromDomainToDB(participant entities.Participant) *pgstore.Participant {
	return &pgstore.Participant{
		ID:          participant.Id.UUID,
		TripID:      participant.TripId.UUID,
		Email:       participant.Email.Value,
		IsConfirmed: participant.ConfirmedPresence,
	}
}
