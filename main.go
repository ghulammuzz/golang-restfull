package main

import (
	"ghulammuzz/golang-restfull/app"
	"ghulammuzz/golang-restfull/controller"
	"ghulammuzz/golang-restfull/helper"
	"ghulammuzz/golang-restfull/repository"
	"ghulammuzz/golang-restfull/service"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
)

func main() {


	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	
	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr: "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}