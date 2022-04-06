package service

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/zogyi/test2/models"
)

type LatestNProjectsFolksRequest struct {
	LatestN uint `json:"latestN"`
}

type LatestNProjectsFolksResponse struct {
	ErrMsg  string               `json:"errMsg"`
	Success bool                 `json:"success"`
	Entity  models.ProjectsFolks `json:"entity"`
}

//the endpoint will receive a request, convert to the request to uint as the lastN param
//it will return the lastN's project names(comma separated) and the sum of folks count
func makeLatestNProjectsFolksEndpoint(service IProjectService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		reqObj := request.(LatestNProjectsFolksRequest)
		names, folks, err := service.GetLastNProjectsFolks(ctx, reqObj.LatestN)
		if err != nil {
			return LatestNProjectsFolksResponse{
				Success: true,
				Entity: models.ProjectsFolks{
					Projects:   names,
					FolksCount: folks},
			}, nil
		}
		return LatestNProjectsFolksResponse{Success: false, ErrMsg: err.Error()}, err
	}
}
