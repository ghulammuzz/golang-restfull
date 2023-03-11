package service

import (
	"context"
	"database/sql"
	"ghulammuzz/golang-restfull/exception"
	"ghulammuzz/golang-restfull/helper"
	"ghulammuzz/golang-restfull/model/domain"
	"ghulammuzz/golang-restfull/model/web"
	"ghulammuzz/golang-restfull/repository"

	"github.com/go-playground/validator/v10"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB 				*sql.DB
	validate		*validator.Validate
}

func NewCategoryService(categoryRepository repository.CategoryRepository, DB *sql.DB, validate *validator.Validate) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: categoryRepository,
		DB: DB,
		validate: validate,
	}
}


func (service *CategoryServiceImpl)Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse{
	err :=service.validate.Struct(request)
	helper.PanicIfError(err)

	tx, err :=service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommirOrRollback(tx)

	category := domain.Category{
		Name : request.Name,
	}

	category = service.CategoryRepository.Save(ctx, tx, category)

	return helper.ToCategoryResponse(category)

}
func (service *CategoryServiceImpl)Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse{
	err :=service.validate.Struct(request)
	helper.PanicIfError(err)

	tx, err :=service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommirOrRollback(tx)


	category, err := service.CategoryRepository.FindById(ctx, tx, request.Id)
	
	if err != nil {
		panic(exception.NewNotFound(err.Error()))
	}
	// helper.PanicIfError(err)

	category.Name = request.Name

	category = service.CategoryRepository.Update(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}
func (service *CategoryServiceImpl)Delete(ctx context.Context, categoryId int) {
	tx, err :=service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommirOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	// helper.PanicIfError(err)
	if err != nil {
		panic(exception.NewNotFound(err.Error()))
	}

	service.CategoryRepository.Delete(ctx, tx, category)

}
func (service *CategoryServiceImpl)FindById(ctx context.Context, categoryId int) web.CategoryResponse{
	tx, err :=service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommirOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	// helper.PanicIfError(err)
	if err != nil {
		panic(exception.NewNotFound(err.Error()))
	}

	return helper.ToCategoryResponse(category)

}
func (service *CategoryServiceImpl)FindAll(ctx context.Context) []web.CategoryResponse{
	tx, err :=service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommirOrRollback(tx)

	categories := service.CategoryRepository.FindAll(ctx, tx)

	
	return helper.ToCategoryResponses(categories)
}