package honeycombs

import "time"

type Turn struct {
	ID        uint
	UserID    uint
	GameID    uint
	CreatedAt time.Time
	Points    uint8
}
