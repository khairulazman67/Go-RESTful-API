package main

import (
	"net/http"
	"resfull_api/app"
	"resfull_api/controller"
	"resfull_api/helper"
	"resfull_api/middleware"
	"resfull_api/repository"
	"resfull_api/service"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
)

func main() {
	db:=app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)
	
	server := http.Server{
		Addr: "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicError(err)

}