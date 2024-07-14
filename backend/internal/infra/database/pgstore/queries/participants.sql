-- name: GetParticipant :one
SELECT
    "id", "trip_id", "email", "is_confirmed"
FROM participants
WHERE
    id = $1;

-- name: ConfirmParticipant :exec
SELECT
    "id", "trip_id", "email", "is_confirmed"
FROM participants
WHERE
    id = $1;


-- name: GetParticipants :many
SELECT
    "id", "trip_id", "email", "is_confirmed"
FROM participants
WHERE
    trip_id = $1;

-- name: InviteParticipantsToTrip :copyfrom
INSERT INTO participants
    ( "trip_id", "email" ) VALUES
    ( $1, $2 );

