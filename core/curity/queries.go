package curity

import (
	"workspace-service/graphql"
)

type AccountFields struct {
	Id          string
	DisplayName string
	Active      bool
	Emails      []struct {
		Value   string
		Primary bool
		//type string// TODO: conflict with go
		Display string
	}
	Meta struct {
		Created      int64
		LastModified int64
		TimeZoneId   string
		ResourceType string
	}
}

func (account *AccountFields) ToDto() *graphql.User {
	return &graphql.User{
		Name: account.DisplayName,
	}
}

type GetAccounts struct {
	Accounts struct {
		TotalCount int32
		PageInfo   struct {
			EndCursor   string
			HasNextPage bool
		}
		Edges []struct {
			Node AccountFields
		}
	} `graphql:"accounts(activeAccountsOnly: $activeOnly, first: $first, after: $after)"`
}

type GetAccountById struct {
	AccountById AccountFields `graphql:"accountById(accountId: $accountId)"`
}

type CreateAccount struct {
	DeleteAccountById struct {
		account AccountFields
	} `graphql:"createAccount(input: {fields: $fields})"`
}

type UserFields struct {
}

type CreateAccountInput struct {
	Fields UserFields
}

type UpdateAccountById struct {
	UpdateAccountById struct {
		account AccountFields
	} `graphql:"updateAccountById(input: {accountId: $accountId, fields: $fields})"`
}

type UpdateAccountByIdInput struct {
	AccountId string
	Fields    UserFields
}

type DeleteAccountById struct {
	DeleteAccountById struct {
		deleted bool
		account AccountFields
	} `graphql:"deleteAccountById(input: {accountId: $accountId})"`
}

type DeleteAccountByIdInput struct {
	AccountId string
}
