// Package ToDoListService GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package ToDoListService

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/create_task": {
            "post": {
                "description": "Create a task and add it to the todo list.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "root"
                ],
                "parameters": [
                    {
                        "description": "Task's title and description",
                        "name": "task",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/task.TaskData"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Created a task successfully",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/get_all_tasks": {
            "get": {
                "description": "List all the items in the todo list.",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "root"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/task.TaskData"
                            }
                        }
                    }
                }
            }
        },
        "/remove_task/{id}": {
            "get": {
                "description": "Remove a task from the todo list.",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "root"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "uuid",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Removed task {id} successfully",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "task.TaskData": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}