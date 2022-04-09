package model

type CategoryCreateRequest struct {
	Name string `json:"name"`
}

type CategoryUpdateRequest struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type CategoryResponse struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}