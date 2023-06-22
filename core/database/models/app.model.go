package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type App struct {
	ID               uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Name             string
	Url              string
	Position         int32
	AvailableLocales []string `gorm:"type:text[]"`
	ExpiresAt        time.Time
}
