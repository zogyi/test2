package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	models "github.com/zogyi/test2/models"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	queryURL      = `https://gitlab.com/api/graphql`
	contentType   = `application/json`
	query         = `query last_projects($n: Int = %d) {projects(last:$n) {nodes {name description forksCount}}}`
	operationName = `last_projects`
)

//IProjectRepository the repository interface for the project
type IProjectRepository interface {
	GetLatestNProjects(ctx context.Context, count uint) (responseObj models.ProjectResponse, err error)
}

//ProjectRepository the repository implement for the project
type ProjectRepository struct{}

//GetLatestNProjects get latest N's projects
func (repository *ProjectRepository) GetLatestNProjects(ctx context.Context, count uint) (responseObj models.ProjectResponse, err error) {
	type queryObj struct {
		Query         string      `json:"query"`
		variables     interface{} `json:"variables"`
		OperationName string      `json:"operationName"`
	}
	var (
		queryByte, responseByte []byte
		response                *http.Response
		responseStr             string
	)

	currentQuery := fmt.Sprintf(query, count)
	query := queryObj{Query: currentQuery, OperationName: operationName}
	if queryByte, err = json.Marshal(query); err != nil {
		return
	}
	if response, err = http.Post(queryURL, contentType, bytes.NewBuffer(queryByte)); err != nil {
		return
	}
	if response.StatusCode == 200 {
		defer response.Body.Close()
		if responseByte, err = ioutil.ReadAll(response.Body); err != nil {
			return
		}
		responseStr = string(responseByte)
		if strings.Contains(responseStr, `errors`) {
			return responseObj, errors.New(`server internal error, please contact the system admin`)
		}
		err = json.Unmarshal(responseByte, &responseObj)
		return
	} else {
		return responseObj, errors.New(`server internal error`)
	}
}
