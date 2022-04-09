package service

import (
	"context"

	"github.com/muhangga/model"
)

type CategoryService interface {
	Create(ctx context.Context, request model.CategoryCreateRequest) model.CategoryResponse
	Update(ctx context.Context, request model.CategoryUpdateRequest) model.CategoryResponse
	Delete(ctx context.Context, categoryId int)
	FindById(ctx context.Context, categoryId int) model.CategoryResponse
	FindAll(ctx context.Context) []model.CategoryResponse
}