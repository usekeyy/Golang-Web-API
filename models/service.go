package models

type Service interface {
	FindAll() ([]Task, error)
	FindById(ID int) (Task, error)
	Create(taskRequest TaskRequest) (Task, error)
	Update(ID int, taskRequest TaskRequest) (Task, error)
	Delete(ID int) (Task, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Task, error) {
	task, err := s.repository.FindAll()
	return task, err
}

func (s *service) FindById(ID int) (Task, error) {
	task, err := s.repository.FindById(ID)
	return task, err
}

func (s *service) Create(taskRequest TaskRequest) (Task, error) {
	task := Task{
		Task_detail: taskRequest.Task_detail,
		Assigne:     taskRequest.Assigne,
		Deadline:    taskRequest.Deadline,
	}
	task, err := s.repository.Create(task)
	return task, err
}

func (s *service) Update(ID int, taskRequest TaskRequest) (Task, error) {
	task, _ := s.repository.FindById(ID)

	task.Task_detail = taskRequest.Task_detail
	task.Assigne = taskRequest.Assigne
	task.Deadline = taskRequest.Deadline

	newTask, err := s.repository.Update(task)
	return newTask, err
}

func (s *service) Delete(ID int) (Task, error) {
	task, _ := s.repository.FindById(ID)

	newTask, err := s.repository.Delete(task)
	return newTask, err
}
