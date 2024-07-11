package api

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/JP-Go/wilson/backend/internal/application/usecases"
	"github.com/JP-Go/wilson/backend/internal/domain/entities"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type API struct {
	confirmParticipantUseCase usecases.ConfirmParticipantPresenceUseCase
	getParticipantUseCase     usecases.GetParticipantUseCase
	logger                    slog.Logger
}

// Confirms a participant on a trip.
// (PATCH /participants/{participantId}/confirm)
func (ap *API) PatchParticipantsParticipantIDConfirm(w http.ResponseWriter, r *http.Request, participantID string) *Response {
	id, err := uuid.Parse(participantID)
	if err != nil {
		return PatchParticipantsParticipantIDConfirmJSON400Response(Error{
			Message: "Invalid UUID",
		})
	}

	participant, err := ap.getParticipantUseCase.Execute(r.Context(), entities.Participant{
		Id: uuid.NullUUID{
			UUID:  id,
			Valid: true,
		},
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return PatchParticipantsParticipantIDConfirmJSON400Response(Error{
				Message: "Error: " + err.Error(),
			})
		}
		message := fmt.Sprintf(`Unhandled error found. Failed to get participant! {
                ctx: PatchParticipantsParticipantIDConfirm,
                parameters: {
                    writter: %v,
                    request: %v,
                    params: {
                        participantID: %v
                    }
                },
                timestamp: %v UTC
            }`,
			w, r, participantID,
			time.Now().UTC(),
		)
		ap.logger.Log(r.Context(), slog.LevelError.Level(), message)
		return PatchParticipantsParticipantIDConfirmJSON400Response(Error{
			Message: "Error: something went wrong",
		})
	}

	if participant.ConfirmedPresence {
		return PatchParticipantsParticipantIDConfirmJSON400Response(Error{
			Message: "Error: participant already confirmed presence",
		})
	}

	if err = ap.confirmParticipantUseCase.Execute(r.Context(), participant); err != nil {
		message := fmt.Sprintf(`Unhandled error found. Failed to confirm participant presence! {
                ctx: PatchParticipantsParticipantIDConfirm,
                parameters: {
                    writter: %v,
                    request: %v,
                    params: {
                        participantID: %v
                    }
                },
                timestamp: %v UTC
            }`,
			w, r, participantID,
			time.Now().UTC(),
		)
		ap.logger.Log(r.Context(), slog.LevelError.Level(), message)
		return PatchParticipantsParticipantIDConfirmJSON400Response(Error{
			Message: "Error: something went wrong",
		})
	}
	return PatchParticipantsParticipantIDConfirmJSON204Response(nil)
}

// Create a new trip
// (POST /trips)
func (api *API) PostTrips(w http.ResponseWriter, r *http.Request) *Response {
	panic("Not implemented")
}

// Get a trip details.
// (GET /trips/{tripId})
func (api *API) GetTripsTripID(w http.ResponseWriter, r *http.Request, tripID string) *Response {
	panic("Not implemented")
}

// Update a trip.
// (PUT /trips/{tripId})
func (api *API) PutTripsTripID(w http.ResponseWriter, r *http.Request, tripID string) *Response {
	panic("Not implemented")
}

// Get a trip activities.
// (GET /trips/{tripId}/activities)
func (api *API) GetTripsTripIDActivities(w http.ResponseWriter, r *http.Request, tripID string) *Response {
	panic("Not implemented")
}

// Create a trip activity.
// (POST /trips/{tripId}/activities)
func (api *API) PostTripsTripIDActivities(w http.ResponseWriter, r *http.Request, tripID string) *Response {
	panic("Not implemented")
}

// Confirm a trip and send e-mail invitations.
// (GET /trips/{tripId}/confirm)
func (api *API) GetTripsTripIDConfirm(w http.ResponseWriter, r *http.Request, tripID string) *Response {
	panic("Not implemented")
}

// Invite someone to the trip.
// (POST /trips/{tripId}/invites)
func (api *API) PostTripsTripIDInvites(w http.ResponseWriter, r *http.Request, tripID string) *Response {
	panic("Not implemented")
}

// Get a trip links.
// (GET /trips/{tripId}/links)
func (api *API) GetTripsTripIDLinks(w http.ResponseWriter, r *http.Request, tripID string) *Response {
	panic("Not implemented")
}

// Create a trip link.
// (POST /trips/{tripId}/links)
func (api *API) PostTripsTripIDLinks(w http.ResponseWriter, r *http.Request, tripID string) *Response {
	panic("Not implemented")
}

// Get a trip participants.
// (GET /trips/{tripId}/participants)
func (api *API) GetTripsTripIDParticipants(w http.ResponseWriter, r *http.Request, tripID string) *Response {
	panic("Not implemented")
}
