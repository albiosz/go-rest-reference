package honeycombs

import "time"

type UserGameState string

const (
	ACTIVE   UserGameState = "ACTIVE"
	FINISHED UserGameState = "FINISHED"
)

type UserGame struct {
	UserID    uint
	GameID    uint
	CreatedAt time.Time
	PlayerNo  uint8
	State     UserGameState
}
