package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var ErrTokenExpired = errors.New("token is expired")
var ErrInvalidToken = errors.New("invalid token")

type Payload struct {
	ID         uuid.UUID `json:"id"`
	Email      string    `json:"email"`
	IssuedTime time.Time `json:"issued_time"`
	ExpireTime time.Time `json:"expire_time"`
}

func NewPayload(email string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:         tokenID,
		Email:      email,
		IssuedTime: time.Now(),
		ExpireTime: time.Now().Add(duration),
	}

	return payload, nil
}

func (p *Payload) Valid() error {
	if time.Now().After(p.ExpireTime) {
		return ErrTokenExpired
	}
	return nil
}
