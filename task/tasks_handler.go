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
	// generate a uuid and add the pair (id, task) to the tasks map
	id := uuid.New().String()
	taskHandler.AddTask(task, id)
	return id
}

func (taskHandler *TasksHandler) AddTask(task *TaskData, id string) {
	// add the pair (id, task) to the tasks map
	taskHandler.tasks[id] = task
}

func (taskHandler *TasksHandler) GetAllTasks() map[string]*TaskData {
	// return all tasks & ids
	return taskHandler.tasks
}

func (taskHandler *TasksHandler) RemoveTask(taskId string) error {
	// remove the task related to the taskId recieved

	if _, ok := taskHandler.tasks[taskId]; !ok {
		return fmt.Errorf("RemoveTask failed, no such task with id: %v", taskId)
	}

	delete(taskHandler.tasks, taskId)

	return nil
}
