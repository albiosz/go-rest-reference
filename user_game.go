package honeycombs

import "time"

type UserGameState string

const (
	ACTIVE   UserGameState = "ACTIVE"
	FINISHED UserGameState = "FINISHED"
)

type UserGame struct {
	UserEmail string
	GameID    uint8
	CreatedAt time.Time
	PlayerNo  uint8
	State     UserGameState
}
