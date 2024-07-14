package repository

import (
	"context"

	"github.com/JP-Go/wilson/backend/internal/domain/entities"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ParticipantRepository interface {
	GetParticipant(ctx context.Context, participant entities.Participant) (*entities.Participant, error)
	ConfirmParticipantPresence(ctx context.Context, participant entities.Participant) (*entities.Participant, error)
}

type PgxPoolParticipantRepository struct {
	pool *pgxpool.Pool
}

// ConfirmParticipantPresence implements ParticipantRepository.
func (p PgxPoolParticipantRepository) ConfirmParticipantPresence(ctx context.Context, participant entities.Participant) (*entities.Participant, error) {
	panic("unimplemented")
}

// GetParticipant implements ParticipantRepository.
func (p PgxPoolParticipantRepository) GetParticipant(ctx context.Context, participant entities.Participant) (*entities.Participant, error) {
	panic("unimplemented")
}

func NewParticipantRepository(pool *pgxpool.Pool) ParticipantRepository {
	return PgxPoolParticipantRepository{
		pool: pool,
	}
}
