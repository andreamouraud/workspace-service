package utils

import (
	"fmt"
	"workspace-service/core/database/models"
	"workspace-service/graphql"

	"github.com/graph-gophers/dataloader"
)

func MapDataloaderResponse(keys dataloader.Keys, users []models.User) []*dataloader.Result {
	userById := map[string]*graphql.User{}
	for _, user := range users {
		userById[user.ID.String()] = user.ToDto()
	}

	// return users in the same order requested
	output := make([]*dataloader.Result, len(keys))
	for index, userKey := range keys {
		user, ok := userById[userKey.String()]
		if ok {
			output[index] = &dataloader.Result{Data: user, Error: nil}
		} else {
			err := fmt.Errorf("user not found %s", userKey.String())
			output[index] = &dataloader.Result{Data: nil, Error: err}
		}
	}
	return output
}
