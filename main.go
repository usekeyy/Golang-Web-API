package main

import (
	controllers "Go-App/controllers"
	"Go-App/models"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	//connection to database
	dsn := "root:@tcp(127.0.0.1:3306)/go_app?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Connection to database error")
	}
	//table migration
	db.AutoMigrate(&models.Task{})

	//new repository
	taskRepository := models.NewRepository(db)

	taskService := models.NewService(taskRepository)

	taskController := controllers.TaskController(taskService)

	//initial router
	router := gin.Default()

	router.GET("/task", taskController.GetAll)
	router.GET("/task/:id", taskController.GetTask)
	router.POST("/task", taskController.Store)
	router.PUT("/task/:id", taskController.Update)
	router.DELETE("/task/:id", taskController.Delete)

	router.Run()
}
