package graphql

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import (
	"workspace-service/core/services"

	graphql_client "github.com/hasura/go-graphql-client"
	"gorm.io/gorm"
)

type Resolver struct {
	userService  services.IUserService
	groupService services.IGroupService
	appService   services.IAppService
}

func NewResolvers(db *gorm.DB, curityClient *graphql_client.Client) *Resolver {
	return &Resolver{
		userService:  services.NewUserService(db, curityClient),
		groupService: services.NewGroupService(db),
		appService:   services.NewAppService(db),
	}
}
