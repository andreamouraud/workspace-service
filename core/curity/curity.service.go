package curity

import (
	"context"
	"workspace-service/graphql"

	graphql_client "github.com/hasura/go-graphql-client"
)

type ICurityService interface {
	CreateAccount()
	UpdateAccountById(string)
	GetAccountById(string)
	GetAccounts()
	DeleteAccountById(string)
}

type CurityService struct {
	curityClient *graphql_client.Client
}

func NewCurityService(curityClient *graphql_client.Client) *CurityService {
	service := &CurityService{curityClient}
	return service
}

func (s *CurityService) GetAccountById(id string) (*AccountFields, error) {
	variables := map[string]interface{}{
		"accountId": id,
	}

	var res GetAccountById
	err := s.curityClient.Query(context.Background(), &res, variables)
	if err != nil {
		return nil, err
	}
	return &res.AccountById, nil
}

func (s *CurityService) GetAccounts() ([]*AccountFields, error) {
	variables := map[string]interface{}{
		"activeOnly": true,
	}

	var res GetAccounts
	err := s.curityClient.Query(context.Background(), &res, variables)
	if err != nil {
		return nil, err
	}

	users := make([]*AccountFields, len(res.Accounts.Edges))
	for i, element := range res.Accounts.Edges {
		copy := element.Node
		users[i] = &copy
	}
	return users, nil
}

func (s *CurityService) DeleteAccountById(id string) error {
	variables := map[string]interface{}{
		"input": DeleteAccountByIdInput{
			AccountId: id,
		},
	}

	var res DeleteAccountById
	return s.curityClient.Query(context.Background(), &res, variables)
}

func (s *CurityService) CreateAccount(input graphql.CreateUserInput) error {
	variables := map[string]interface{}{
		"input": UserFields{},
	}

	var res CreateAccount
	return s.curityClient.Query(context.Background(), &res, variables)
}

func (s *CurityService) UpdateAccountById(input graphql.UpdateUserInput) error {
	variables := map[string]interface{}{
		"input": UpdateAccountByIdInput{},
	}

	var res CreateAccount
	return s.curityClient.Query(context.Background(), &res, variables)
}
