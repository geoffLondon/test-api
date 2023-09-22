package utils

import "github.com/google/uuid"

type UUIDGenerator interface {
	New() string
}

type UUIDGeneratorImpl struct{}

func (r *UUIDGeneratorImpl) New() string {
	return uuid.New().String()
}
