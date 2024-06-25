package utils

import "github.com/google/uuid"

func CreateId() string {
	id := uuid.New().String()
	return id
}
