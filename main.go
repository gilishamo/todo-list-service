package main

import (
	"fmt"
	"os"

	_ "todo-list-service/docs/ToDoListService"
	service "todo-list-service/todo_list_server"
)

//	@title			Go + Echo Todo List API
//	@version		1.0
//	@description	This is a simple todo list server.

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @host		localhost:8080
// @BasePath	/
// @schemes	http
func main() {
	args := os.Args

	if len(args) < 3 {
		fmt.Printf("ERROR: you must provide an ip")
	}

	serviceAddress := args[1]
	var servicesAddress []string
	if len(args) > 2 {
		servicesAddress = args[2:]
	}

	// initialize service & start
	s := service.NewToDoListServer(serviceAddress, servicesAddress)
	s.Start()
}
