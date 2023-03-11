package service

import (
	"context"
	"ghulammuzz/golang-restfull/model/web"
)

// representasi request
type CategoryService interface {
	Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse
	Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse
	Delete(ctx context.Context, categoryId int)
	FindById(ctx context.Context,categoryId int) web.CategoryResponse
	FindAll(ctx context.Context) []web.CategoryResponse
}