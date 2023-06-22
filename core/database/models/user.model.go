package models

import (
	"workspace-service/graphql"

	"github.com/gofrs/uuid"
)

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Name     string
	CurityID uuid.UUID
	Groups   []*Group `gorm:"many2many:user_group;"`
}

func (user *User) ToDto() *graphql.User {
	return &graphql.User{
		ID:   user.ID.String(),
		Name: "Toto",
	}
}
