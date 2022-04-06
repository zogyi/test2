package service

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/zogyi/test2/models"
)

type LatestNProjectsFolksReq struct {
	LatestN uint `json:"latestN"`
}
type ProjectResponse struct {
}

//the endpoint will receive a request, convert to the request to uint as the lastN param
//it will return the lastN's project names(comma separated) and the sum of folks count
func makeGetLatestProjectsFolksEndpoint(service IProjectService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		n := request.(uint)
		names, folks, err := service.GetLastNProjectsFolks(ctx, n)
		if err != nil {
			return models.ResponseResult{
				Success: true,
				Entity: models.ProjectsFolks{
					Projects:   names,
					FolksCount: folks},
			}, nil
		}
		return models.ResponseResult{Success: false, ErrMsg: err.Error()}, err
	}
}
