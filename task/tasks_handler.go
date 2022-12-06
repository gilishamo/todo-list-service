package task

import (
	"fmt"

	"github.com/google/uuid"
)

type TasksHandler struct {
	tasks map[string]*TaskData
}

func NewTasksHandler() *TasksHandler {
	return &TasksHandler{
		tasks: make(map[string]*TaskData, 0),
	}
}

func (taskHandler *TasksHandler) CreateTask(task *TaskData) string {
	id := uuid.New().String()
	taskHandler.tasks[id] = task
	return id
}

func (taskHandler *TasksHandler) AddTask(task *TaskData, id string) {
	taskHandler.tasks[id] = task
}

func (taskHandler *TasksHandler) GetAllTasks() map[string]*TaskData {
	return taskHandler.tasks
}

func (taskHandler *TasksHandler) RemoveTask(taskId string) error {
	if _, ok := taskHandler.tasks[taskId]; !ok {
		return fmt.Errorf("RemoveTask failed, no such task with id: %v", taskId)
	}

	delete(taskHandler.tasks, taskId)

	return nil
}
