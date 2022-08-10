package controllers

import (
	"Go-App/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type taskController struct {
	taskService models.Service
}

func TaskController(taskService models.Service) *taskController {
	return &taskController{taskService}
}

func (h *taskController) GetAll(c *gin.Context) {
	task, err := h.taskService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	//mapping response
	var response []models.TaskResponse

	for _, t := range task {
		taskResponse := convertToTaskResponse(t)
		response = append(response, taskResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": response,
	})
}

func (h *taskController) GetTask(c *gin.Context) {
	requestId := c.Param("id")
	id, _ := strconv.Atoi(requestId)
	task, err := h.taskService.FindById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	response := convertToTaskResponse(task)
	c.JSON(http.StatusOK, gin.H{
		"data": response,
	})
}

func (h *taskController) Store(c *gin.Context) {
	var taskRequest models.TaskRequest

	err := c.ShouldBindJSON(&taskRequest)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			message := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, message)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}
	task, err := h.taskService.Create(taskRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Task Successfully Created",
		"data":    task,
	})
}

func (h *taskController) Update(c *gin.Context) {
	var taskRequest models.TaskRequest

	err := c.ShouldBindJSON(&taskRequest)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			message := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, message)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	requestId := c.Param("id")
	id, _ := strconv.Atoi(requestId)
	task, err := h.taskService.Update(id, taskRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Task Successfully Updated",
		"data":    task,
	})
}

func (h *taskController) Delete(c *gin.Context) {
	requestId := c.Param("id")
	id, _ := strconv.Atoi(requestId)
	task, err := h.taskService.Delete(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Task Successfully Deleted",
		"data":    task,
	})
}

func convertToTaskResponse(task models.Task) models.TaskResponse {
	response := models.TaskResponse{
		ID:          task.ID,
		Task_detail: task.Task_detail,
		Assigne:     task.Assigne,
		Deadline:    task.Deadline,
	}
	return response
}

func RootController(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, World!",
	})
}
