package models

import (
	"github.com/gofrs/uuid"
)

type Group struct {
	ID    uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Name  string    `form:"first_name" json:"first_name,omitempty"`
	Users []*User   `gorm:"many2many:user_group;"`
}
