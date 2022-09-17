package helper

import "github.com/google/uuid"

func CreateNewUUID() string {
	uuid := uuid.New()
	return uuid.String()
}
