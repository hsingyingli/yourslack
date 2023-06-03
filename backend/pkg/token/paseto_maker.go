package token

import (
	"fmt"
	"time"

	db "github.com/hsingyingli/yourslack-backend/db/sqlc"
	"github.com/o1egl/paseto"
)

type PasetoMaker struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

func NewPasetoMaker(symmetricKey string) (Maker, error) {
	if len(symmetricKey) != 32 {
		return nil, fmt.Errorf("length of symmetricKey must be greater than 32, but got: %v", len(symmetricKey))
	}

	maker := &PasetoMaker{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
	}
	return maker, nil
}

func (pasetoMaker PasetoMaker) CreateToken(user db.User, duration time.Duration) (string, *Payload, error) {
	payload, err := NewPayload(user, duration)
	if err != nil {
		return "", nil, err
	}
	token, err := pasetoMaker.paseto.Encrypt(pasetoMaker.symmetricKey, payload, nil)
	return token, payload, err
}

func (pasetoMaker PasetoMaker) VerifyToken(token string) (*Payload, error) {
	var payload Payload
	err := pasetoMaker.paseto.Decrypt(token, pasetoMaker.symmetricKey, &payload, nil)
	if err != nil {
		return nil, err
	}

	err = payload.Vaild()

	if err != nil {
		return nil, err
	}

	return &payload, nil
}
