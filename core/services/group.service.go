package services

import (
	"workspace-service/core/database/models"
	"workspace-service/graphql"
	"workspace-service/utils"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type IGroupService interface {
	GetGroups() ([]*graphql.Group, error)
	GetGroup(uuid.UUID) (*graphql.Group, error)
	CreateGroup(*graphql.CreateGroupInput) (*graphql.Group, error)
	UpdateGroup(uuid.UUID, *graphql.UpdateGroupInput) (*graphql.Group, error)
	DeleteGroup(uuid.UUID) (bool, error)
}

type GroupService struct {
	db *gorm.DB
}

func NewGroupService(db *gorm.DB) *GroupService {
	service := &GroupService{db}
	return service
}

func (s *GroupService) GetGroups() ([]*graphql.Group, error) {
	groups := []*models.Group{}
	err := s.db.Find(&groups).Error
	return utils.Map(groups, s.buildDto), err
}

func (s *GroupService) GetGroup(id uuid.UUID) (*graphql.Group, error) {
	group := &models.Group{}
	err := s.db.Where("id = ?", id).First(group).Error
	return s.buildDto(group), err
}

func (s *GroupService) CreateGroup(input *graphql.CreateGroupInput) (*graphql.Group, error) {
	group := &models.Group{
		Name:  input.Name,
		Users: []*models.User{},
	}
	err := s.db.Create(&group).Error
	return s.buildDto(group), err
}

func (s *GroupService) UpdateGroup(id uuid.UUID, input *graphql.UpdateGroupInput) (*graphql.Group, error) {
	group := &models.Group{
		ID:    id,
		Name:  *input.Name,
		Users: []*models.User{},
	}
	err := s.db.Model(&group).Where("id = ?", id).Updates(group).Error
	return s.buildDto(group), err
}

func (s *GroupService) DeleteGroup(id uuid.UUID) (bool, error) {
	group := &models.Group{}
	err := s.db.Delete(group, id).Error
	hasError := err != nil
	return hasError, err
}

func (s *GroupService) buildDto(group *models.Group) *graphql.Group {
	return &graphql.Group{
		ID:   group.ID.String(),
		Name: group.Name,
	}
}
