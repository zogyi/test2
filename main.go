package main

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/zogyi/test2/repository"
	"github.com/zogyi/test2/service"
	"os"
)

func main() {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	projectService := service.NewIProjectService(&repository.ProjectRepository{}, logger)
	projectNames, folksCount, err := projectService.GetLastNProjectsFolks(context.Background(), 10)
	fmt.Println(projectNames)
	fmt.Println(folksCount)
	fmt.Println(err)
}
