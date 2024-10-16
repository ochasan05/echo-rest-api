package main

import (
	"go-rest-api-todo/controller"
	"go-rest-api-todo/db"
	"go-rest-api-todo/repository"
	"go-rest-api-todo/router"
	"go-rest-api-todo/usecase"
	"go-rest-api-todo/validator"
)

func main() {
	db := db.NewDB()
	userRepository := repository.NewUserRepository(db)
	userValidator := validator.NewUserValidator()
	userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	userController := controller.NewUserController(userUsecase)

	taskRepository := repository.NewTaskRepository(db)
	taskValidator := validator.NewTaskValidator()
	taskUsecase := usecase.NewTaskUsecase(taskRepository, taskValidator)
	taskController := controller.NewTaskController(taskUsecase)

	e := router.NewRouter(userController, taskController)
	e.Logger.Fatal(e.Start(":80"))
}
