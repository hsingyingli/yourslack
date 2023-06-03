package token

import (
	"time"

	db "github.com/hsingyingli/yourslack-backend/db/sqlc"
)

type Maker interface {
	CreateToken(user db.User, duration time.Duration) (string, *Payload, error)
	VerifyToken(token string) (*Payload, error)
}
