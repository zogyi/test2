package service

import (
	"context"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/zogyi/test2/models"
	repository "github.com/zogyi/test2/repository"
	"strings"
)

type IProjectService interface {
	GetLastNProjectsFolks(ctx context.Context, count uint) (string, uint, error)
}

func NewIProjectService(repository repository.IProjectRepository, logger log.Logger) IProjectService {
	return &projectService{
		IProjectRepository: repository,
		logger:             logger,
	}
}

type projectService struct {
	logger log.Logger
	repository.IProjectRepository
}

func (service *projectService) GetLastNProjectsFolks(ctx context.Context, count uint) (names string, totalCount uint, err error) {
	var (
		responseObj models.ProjectResponse
		nameArray   []string
	)
	if responseObj, err = service.GetLatestNProjects(ctx, count); err != nil {
		level.Error(service.logger).Log(`message`, `can't get the service`)
		return
	}
	for _, node := range responseObj.Data.Projects.Nodes {
		totalCount = totalCount + node.ForksCount
		nameArray = append(nameArray, node.Name)
	}
	level.Info(service.logger).Log(`message`, `can't get the service`)
	return strings.Join(nameArray, `,`), totalCount, nil
}
