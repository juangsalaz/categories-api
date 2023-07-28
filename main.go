package main

import (
	"juang/belajar-golang-restful-api/app"
	"juang/belajar-golang-restful-api/controller"
	"juang/belajar-golang-restful-api/exception"
	"juang/belajar-golang-restful-api/helper"
	"juang/belajar-golang-restful-api/middleware"
	"juang/belajar-golang-restful-api/repository"
	"juang/belajar-golang-restful-api/service"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func main() {
	db := app.NewDB()
	validate := validator.New()

	categpryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categpryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "127.0.0.1:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
