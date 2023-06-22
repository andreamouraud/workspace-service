package services

import (
	"workspace-service/core/database/models"
	"workspace-service/graphql"
	"workspace-service/utils"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type IAppService interface {
	GetApps() ([]*graphql.App, error)
	GetApp(uuid.UUID) (*graphql.App, error)
	CreateApp(*graphql.CreateAppInput) (*graphql.App, error)
	UpdateApp(uuid.UUID, *graphql.UpdateAppInput) (*graphql.App, error)
	DeleteApp(uuid.UUID) (bool, error)
}

type AppService struct {
	db *gorm.DB
}

func NewAppService(db *gorm.DB) *AppService {
	service := &AppService{db}
	return service
}

func (s *AppService) GetApps() ([]*graphql.App, error) {
	apps := []*models.App{}
	err := s.db.Find(&apps).Error
	return utils.Map(apps, s.buildDto), err
}

func (s *AppService) GetApp(id uuid.UUID) (*graphql.App, error) {
	app := &models.App{}
	err := s.db.Where("id = ?", id).First(app).Error
	return s.buildDto(app), err
}

func (s *AppService) CreateApp(input *graphql.CreateAppInput) (*graphql.App, error) {
	app := &models.App{
		Name: input.Name,
	}
	err := s.db.Create(&app).Error
	return s.buildDto(app), err
}

func (s *AppService) UpdateApp(id uuid.UUID, input *graphql.UpdateAppInput) (*graphql.App, error) {
	app := &models.App{
		ID:   id,
		Name: *input.Name,
	}
	err := s.db.Model(&app).Where("id = ?", id).Updates(app).Error
	return s.buildDto(app), err
}

func (s *AppService) DeleteApp(id uuid.UUID) (bool, error) {
	err := s.db.Delete(id).Error
	hasError := err != nil
	return hasError, err
}

func (s *AppService) buildDto(app *models.App) *graphql.App {
	return &graphql.App{
		ID:   app.ID.String(),
		Name: app.Name,
	}
}
