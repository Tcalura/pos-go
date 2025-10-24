package entity

import (
	"github.com/google/uuid"
)

type ID string

func NewID() ID {
	return ID(uuid.New().String())
}

func ParseID(s string) (ID, error) {
	id, err := uuid.Parse(s)
	if err != nil {
		return "", err
	}
	return ID(id.String()), nil
}