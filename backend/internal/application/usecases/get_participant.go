package usecases

import (
	"context"

	"github.com/JP-Go/wilson/backend/internal/domain/entities"
	"github.com/JP-Go/wilson/backend/internal/domain/repository"
)

type GetParticipantUseCase struct {
	participantRepository repository.ParticipantRepository
}

func NewGetParticipantUseCase(participantRepository repository.ParticipantRepository) GetParticipantUseCase {
	return GetParticipantUseCase{
		participantRepository: participantRepository,
	}
}

func (cpuc *GetParticipantUseCase) Execute(ctx context.Context, participant entities.Participant) (*entities.Participant, error) {

	found, err := cpuc.participantRepository.GetParticipant(ctx, participant)
	if err != nil {
		return nil, err
	}
	return found, nil
}
