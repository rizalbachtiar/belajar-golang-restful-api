package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	"github.com/rizalbachtiar/belajar-golang-restful-api/app"
	"github.com/rizalbachtiar/belajar-golang-restful-api/controller"
	"github.com/rizalbachtiar/belajar-golang-restful-api/exception"
	"github.com/rizalbachtiar/belajar-golang-restful-api/helper"
	"github.com/rizalbachtiar/belajar-golang-restful-api/middleware"
	"github.com/rizalbachtiar/belajar-golang-restful-api/repository"
	"github.com/rizalbachtiar/belajar-golang-restful-api/service"
)

func main() {

	db := app.NewDB()
	validate := validator.New()

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
