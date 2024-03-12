package honeycombs

import "time"

type GameState string

const (
	CREATED     GameState = "CREATED"
	IN_PROGRESS GameState = "IN_PROGRESS"
	ENDED       GameState = "FINISHED"
)

type Game struct {
	ID               uint
	CreatedAt        time.Time
	CreatedByEmail   string
	CreatedBy        *User
	State            GameState
	PlayingUserEmail *string
}
