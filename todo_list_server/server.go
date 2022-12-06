package server

import (
	_ "ToDoListService/docs/ToDoListService"
	"ToDoListService/http_utils"
	"ToDoListService/task"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

const (
	CREATE_TASK   string = "/create_task"
	Get_ALL_TASKS string = "/get_all_tasks"
	Remove_Task   string = "/remove_task"
	INNER_Add     string = "/inner_add_task"
	INNER_Remove  string = "/inner_remove_task"
)

type ToDoListServer struct {
	e               *echo.Echo
	tasksHandler    *task.TasksHandler
	address         string
	servicesAddress []string
}

func NewToDoListServer(serviceAddress string, servicesAddress []string) *ToDoListServer {
	e := echo.New()

	s := new(ToDoListServer)

	// Routes
	e.POST(CREATE_TASK, s.CreateTask)
	e.GET(Get_ALL_TASKS, s.GetAllTasks)
	e.GET(Remove_Task+"/:id", s.RemoveTask)
	e.POST(INNER_Add, s.GetInnerAddRequest)
	e.POST(INNER_Remove, s.GetInnerRemoveRequest)
	// docs route
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	s.e = e
	s.tasksHandler = task.NewTasksHandler()
	s.address = serviceAddress
	s.servicesAddress = servicesAddress

	return s
}

func (server *ToDoListServer) Start() {
	log.Fatal(server.e.Start(server.address))
}

// CreateTask godoc
// @Description Create a task and add it to the todo list.
// @Tags root
// @Accept json
// @Param task body task.TaskData true "Task's title and description"
// @Produce plain
// @Success 200 {string} string "Created a task successfully"
// @Router /create_task [post]
func (server *ToDoListServer) CreateTask(c echo.Context) error {
	log.Print("Recived /create_tasks requets\n")

	var newTask = new(task.TaskData)
	if err := c.Bind(&newTask); err != nil {
		log.Printf("CreateTask, c.Bind failed with error: %v\n", err)
		return c.String(http.StatusBadRequest, err.Error())
	}

	id := server.tasksHandler.CreateTask(newTask)

	go server.SendAddRequestToServers(newTask, id)

	log.Printf("Created a task successfully\n")
	return c.String(http.StatusOK, "Created a task successfully")
}

// GetAllTasks godoc
//
//	@Description	List all the items in the todo list.
//	@Tags			root
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{array} task.TaskData
//	@Router			/get_all_tasks [get]
func (server *ToDoListServer) GetAllTasks(c echo.Context) error {
	log.Printf("Recived /get_all_tasks requets\n")

	return c.JSON(http.StatusOK, server.tasksHandler.GetAllTasks())
}

// RemoveTask godoc
// @Description Remove a task from the todo list.
// @Tags root
// @Accept */*
// @Param id path string true "uuid"
// @Produce plain
// @Success 200 {string} string "Removed task {id} successfully"
// @Router /remove_task/{id} [get]
func (server *ToDoListServer) RemoveTask(c echo.Context) error {
	log.Printf("Recived /remove_task requets\n")

	id := c.Param("id")

	if err := server.tasksHandler.RemoveTask(id); err != nil {
		log.Printf("failed to remove task %v, error is: %v", id, err)
		return c.String(http.StatusBadRequest, fmt.Sprintf("failed to remove task %v, error is: %v", id, err))
	}

	go server.SendRemoveRequestToServers(id)

	log.Printf("removed task %v successfully", id)
	return c.String(http.StatusOK, fmt.Sprintf("Removed task %v successfully", id))
}

func (server *ToDoListServer) SendAddRequestToServers(task *task.TaskData, id string) error {
	log.Printf("Executing SendAddRequestToServers\n")

	params := map[string]string{"id": id}
	headers := map[string]string{"Content-Type": "application/json"}

	return server.SendPostRequestToServer(INNER_Add, task, params, headers)
}

func (server *ToDoListServer) SendRemoveRequestToServers(id string) error {
	log.Printf("Executing SendRemoveRequestToServers\n")

	params := map[string]string{"id": id}
	return server.SendPostRequestToServer(INNER_Remove, nil, params, nil)
}

func (server *ToDoListServer) SendPostRequestToServer(path string, body interface{}, params map[string]string, headers map[string]string) error {
	c := http_utils.NewHttpClient()

	for _, addr := range server.servicesAddress {

		url := fmt.Sprintf("http://%s%s", addr, path)

		data, err := json.Marshal(body)
		if err != nil {
			log.Printf("SendPostRequestToServer, json.Marshal failed with error: %v\n", err)
			return fmt.Errorf("SendPostRequestToServer, json.Marshal failed with error: %v", err)
		}
		body := bytes.NewReader(data)

		res, err := c.SendPostRequest(url, body, params, headers)
		if err != nil {
			log.Printf("SendPostRequestToServer, SendGetRequest failed with error: %v\n", err)
			return fmt.Errorf("SendPostRequestToServer, SendGetRequest failed with error: %v", err)
		}

		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			println(res.Status)
			log.Printf("SendPostRequestToServer response status code is: %v, but expected to be: %v\n", res.StatusCode, http.StatusOK)
			return fmt.Errorf("SendPostRequestToServer response status code is: %v, but expected to be: %v", res.StatusCode, http.StatusOK)
		}

	}

	return nil
}

func (server *ToDoListServer) GetInnerAddRequest(c echo.Context) error {
	log.Printf("Recived /inner_add_task requets\n")

	id := c.QueryParam("id")

	var newTask = new(task.TaskData)
	if err := c.Bind(&newTask); err != nil {
		log.Printf("GetUpdateRequest, c.Bind failed with error: %v\n", err)
		return c.String(http.StatusBadRequest, err.Error())
	}

	server.tasksHandler.AddTask(newTask, id)

	return nil
}

func (server *ToDoListServer) GetInnerRemoveRequest(c echo.Context) error {
	log.Printf("Recived /inner_remove_task requets\n")

	id := c.QueryParam("id")

	if err := server.tasksHandler.RemoveTask(id); err != nil {
		log.Printf("GetInnerRemoveRequest, RemoveTask faled with error: %v\n", err)
		return err
	}

	return nil
}
