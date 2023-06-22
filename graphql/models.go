// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graphql

type App struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type CreateAppInput struct {
	Name string `json:"name"`
}

type CreateGroupInput struct {
	Name string `json:"name"`
}

type CreateUserInput struct {
	Name string `json:"name"`
}

type Group struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type UpdateAppInput struct {
	Name *string `json:"name,omitempty"`
}

type UpdateGroupInput struct {
	Name *string `json:"name,omitempty"`
}

type UpdateUserInput struct {
	Name *string `json:"name,omitempty"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
