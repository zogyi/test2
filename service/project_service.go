package service

import (
	"context"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/zogyi/test2/models"
	repository "github.com/zogyi/test2/repository"
	"strings"
)

//IProjectService the service interface for the project
type IProjectService interface {
	GetLastNProjectsFolks(ctx context.Context, count uint) (string, uint, error)
}

func NewIProjectService(repository repository.IProjectRepository, logger log.Logger) IProjectService {
	return &projectService{
		IProjectRepository: repository,
		logger:             logger,
	}
}

//projectService the service implement for the project
type projectService struct {
	logger log.Logger
	repository.IProjectRepository
}

//GetLastNProjectsFolks get latest N's projects name and folks count
func (service *projectService) GetLastNProjectsFolks(ctx context.Context, count uint) (projects string, folksCount uint, err error) {
	var (
		responseObj models.ProjectResponse
		nameArray   []string
	)
	if responseObj, err = service.GetLatestNProjects(ctx, count); err != nil {
		level.Error(service.logger).Log(`message`, `can't get the service`)
		return
	}
	for _, node := range responseObj.Data.Projects.Nodes {
		folksCount = folksCount + node.ForksCount
		nameArray = append(nameArray, node.Name)
	}
	level.Info(service.logger).Log(`message`, `can't get the service`)
	return strings.Join(nameArray, `,`), folksCount, nil
}
