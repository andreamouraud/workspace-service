package services

import (
	"workspace-service/core/curity"
	"workspace-service/core/database/models"
	"workspace-service/graphql"
	"workspace-service/utils"

	"github.com/gofrs/uuid"
	"github.com/graph-gophers/dataloader"
	graphql_client "github.com/hasura/go-graphql-client"
	"gorm.io/gorm"
)

type IUserService interface {
	GetUsers() ([]*graphql.User, error)
	GetUser(uuid.UUID) (*graphql.User, error)
	CreateUser(*graphql.CreateUserInput) (*graphql.User, error)
	UpdateUser(uuid.UUID, *graphql.UpdateUserInput) (*graphql.User, error)
	DeleteUser(uuid.UUID) (bool, error)
	GetUsersGroups(dataloader.Keys) []*dataloader.Result
}

type UserService struct {
	db            *gorm.DB
	curityService *curity.CurityService
}

func NewUserService(db *gorm.DB, curityClient *graphql_client.Client) *UserService {
	curityService := curity.NewCurityService(curityClient)
	service := &UserService{db, curityService}
	return service
}

func (s *UserService) GetUsers() ([]*graphql.User, error) {
	dbUsers := []*models.User{}
	err := s.db.Find(&dbUsers).Error
	if err != nil {
		return nil, err
	}
	curityUsers, err := s.curityService.GetAccounts()
	if err != nil {
		return nil, err
	}
	curityUsersMap := make(map[string]*graphql.User)
	for i := range curityUsers {
		curityUsersMap[curityUsers[i].Id] = curityUsers[i].ToDto()
	}

	users := []*graphql.User{}
	for i := range dbUsers {
		value, found := curityUsersMap[dbUsers[i].CurityID.String()]
		if found {
			user := dbUsers[i].ToDto()
			user.Name = value.Name
			users = append(users, user)
		}
	}
	return users, err
}

func (s *UserService) GetUser(id uuid.UUID) (*graphql.User, error) {
	dbUser := &models.User{}
	err := s.db.Where("id = ?", id).First(dbUser).Error
	if err != nil {
		return nil, err
	}

	curityUser, err := s.curityService.GetAccountById(id.String())
	if err != nil {
		return nil, err
	}
	user := dbUser.ToDto()
	user.ID = curityUser.Id
	return user, err
}

func (s *UserService) CreateUser(input *graphql.CreateUserInput) (*graphql.User, error) {
	curityUser, err := s.curityService.CreateAccount(input)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		CurityID: uuid.FromStringOrNil(curityUser.Id),
		Name:     input.Name,
	}
	err = s.db.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return s.buildDto(user), err
}

func (s *UserService) UpdateUser(id uuid.UUID, input *graphql.UpdateUserInput) (*graphql.User, error) {
	user := &models.User{
		ID:     id,
		Groups: []*models.Group{},
	}
	err := s.db.Model(&user).Where("id = ?", id).Updates(user).Error
	if err != nil {
		return nil, err
	}

	curityUser, curityErr := s.curityService.UpdateAccountById(id.String(), input)
	if curityErr != nil {
		return nil, curityErr
	}
	user.CurityID = uuid.FromStringOrNil(curityUser.Id)
	return s.buildDto(user), err
}

func (s *UserService) DeleteUser(id uuid.UUID) (bool, error) {
	err := s.db.Delete(id).Error
	if err != nil {
		return false, err
	}
	err = s.curityService.DeleteAccountById(id.String())
	if err != nil {
		return false, err
	}
	return true, nil
}

func (s *UserService) buildDto(user *models.User) *graphql.User {
	return &graphql.User{
		ID:   user.ID.String(),
		Name: "Toto",
	}
}

func (s *UserService) GetUsersGroups(keys dataloader.Keys) []*dataloader.Result {
	// read all requested users in a single query
	userIDs := make([]string, len(keys))
	for ix, key := range keys {
		userIDs[ix] = key.String()
	}
	var users []models.User
	s.db.Raw(
		`SELECT group.* 
		FROM group, user 
		WHERE user_group.user_id IN ?
		AND group.id = user_group.group_id`,
		userIDs,
	).Scan(&users)
	return utils.MapDataloaderResponse(keys, users)
}
