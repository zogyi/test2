package service

import (
	"context"
	"github.com/go-kit/kit/log"
	"github.com/stretchr/testify/assert"
	"github.com/zogyi/test2/repository"
	"os"
	"strings"
	"testing"
)

func TestProjectService_GetLastNProjectsFolks(t *testing.T) {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	assert := assert.New(t)
	repository := repository.ProjectRepository{}
	projectService := NewIProjectService(&repository, logger)
	projectNames, folksCount, err := projectService.GetLastNProjectsFolks(context.Background(), 10)

	assert.Nil(err)
	assert.GreaterOrEqual(folksCount, uint(0), `the total folks count doesn't greater or equal than 0`)
	assert.NotEmpty(projectNames)
	assert.Equal(10, len(strings.Split(projectNames, `,`)), `actual project names count doesn't match expected`)

	projectService = NewIProjectService(&repository, logger)
	projectNames, folksCount, err = projectService.GetLastNProjectsFolks(context.Background(), 0)

	assert.Nil(err)
	assert.GreaterOrEqual(folksCount, uint(0), `the total folks count doesn't greater or equal than 0`)
	assert.Equal(``, projectNames, `actual project names count doesn't match expected`)
}
