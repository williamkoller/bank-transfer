package uuid

import (
	"github.com/google/uuid"
)

type UUID string

func New() UUID {
	return UUID(uuid.NewString())
}

func FromString(id string) (UUID, error) {
	_, err := uuid.Parse(id)
	if err != nil {
		return "", err
	}
	return UUID(id), nil
}
