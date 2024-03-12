package honeycombs

import "time"

type Turn struct {
	ID        uint
	UserEmail string
	GameID    uint
	CreatedAt time.Time
	Points    uint
}
