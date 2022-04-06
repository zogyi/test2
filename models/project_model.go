package models

type Node struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	ForksCount  uint   `json:"forksCount"`
}

type Project struct {
	Nodes []Node `json:"nodes"`
}

type Data struct {
	Projects Project `json:"projects"`
}

type ProjectResponse struct {
	Data Data `json:"data"`
}

type ResponseResult struct {
	ErrMsg  string      `json:"errMsg"`
	Success bool        `json:"success"`
	Entity  interface{} `json:"entity"`
}

type ProjectsFolks struct {
	Projects   string `json:"projects"`
	FolksCount uint   `json:"folksCount"`
}
