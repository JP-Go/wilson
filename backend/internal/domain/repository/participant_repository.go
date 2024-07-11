package repository

import (
	"context"

	"github.com/JP-Go/wilson/backend/internal/domain/entities"
)

type ParticipantRepository interface {
	GetParticipant(ctx context.Context, participant entities.Participant) (*entities.Participant, error)
	ConfirmParticipantPresence(ctx context.Context, participant entities.Participant) (*entities.Participant, error)
}
