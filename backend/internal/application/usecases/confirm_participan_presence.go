package usecases

import (
	"context"

	"github.com/JP-Go/wilson/backend/internal/domain/entities"
	"github.com/JP-Go/wilson/backend/internal/domain/repository"
)

type ConfirmParticipantPresenceUseCase struct {
	participantRepository repository.ParticipantRepository
}

func NewConfirParticipantPresenceUseCase(participan_repository repository.ParticipantRepository) ConfirmParticipantPresenceUseCase {
	return ConfirmParticipantPresenceUseCase{
		participantRepository: participan_repository,
	}
}

func (cpuc *ConfirmParticipantPresenceUseCase) Execute(ctx context.Context, participant *entities.Participant) error {

	_, err := cpuc.participantRepository.ConfirmParticipantPresence(ctx, *participant)

	if err != nil {
		return err
	}
	participant.ConfirmPresence()

	return nil
}
