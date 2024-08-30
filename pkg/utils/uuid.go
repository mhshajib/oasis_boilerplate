package utils

import "github.com/google/uuid"

// GenerateUUID generates a unique Id
func GenerateUUID() string {
	newUUID := uuid.New()
	return newUUID.String()
}
