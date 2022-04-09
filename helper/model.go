package helper

import (
	"github.com/muhangga/entity"
	"github.com/muhangga/model"
)

func ToCategoryResponse(category entity.Category) model.CategoryResponse {
	return model.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func ToCategoryResponses(categories []entity.Category) []model.CategoryResponse {
	var categoryResponses []model.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, ToCategoryResponse(category))
	}
	return categoryResponses
}