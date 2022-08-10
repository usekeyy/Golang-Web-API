package models

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Task, error)
	FindById(ID int) (Task, error)
	Create(task Task) (Task, error)
	Update(task Task) (Task, error)
	Delete(task Task) (Task, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Task, error) {
	var task []Task
	err := r.db.Find(&task).Error

	return task, err
}

func (r *repository) FindById(ID int) (Task, error) {
	var task Task
	err := r.db.Find(&task, ID).Error

	return task, err
}

func (r *repository) Create(task Task) (Task, error) {
	err := r.db.Create(&task).Error

	return task, err
}

func (r *repository) Update(task Task) (Task, error) {
	err := r.db.Save(&task).Error

	return task, err
}

func (r *repository) Delete(task Task) (Task, error) {
	err := r.db.Delete(&task).Error

	return task, err
}
