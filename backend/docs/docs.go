// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

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
        "/admin": {
            "get": {
                "description": "List all admins",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "admin"
                ],
                "summary": "List admins",
                "parameters": [
                    {
                        "type": "string",
                        "description": "full or partial username",
                        "name": "username",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/admin.PagOutput"
                        }
                    }
                }
            },
            "post": {
                "description": "Add a new admin",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "admin"
                ],
                "summary": "Add admin",
                "parameters": [
                    {
                        "description": "Admin input",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/admin.Input"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/admin.Output"
                        }
                    }
                }
            }
        },
        "/admin/{id}": {
            "get": {
                "description": "Get admin by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "admin"
                ],
                "summary": "Get admin",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Admin ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/admin.Output"
                        }
                    }
                }
            },
            "put": {
                "description": "Update admin by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "admin"
                ],
                "summary": "Update admin",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Admin ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Admin update input",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/admin.UpdateInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/admin.Output"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete admin by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "admin"
                ],
                "summary": "Delete admin",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Admin ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    }
                }
            }
        },
        "/auth/login": {
            "post": {
                "description": "Login as an admin",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "description": "login input",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.LoginInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/auth.LoginOutput"
                        }
                    }
                }
            }
        },
        "/barber": {
            "get": {
                "description": "List all barbers",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "barber"
                ],
                "summary": "List barbers",
                "parameters": [
                    {
                        "type": "string",
                        "description": "full or partial barber name",
                        "name": "name",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/barber.PagOutput"
                        }
                    }
                }
            },
            "post": {
                "description": "Add a new barber",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "barber"
                ],
                "summary": "Add barber",
                "parameters": [
                    {
                        "description": "Barber input",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/barber.Input"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/barber.Output"
                        }
                    }
                }
            }
        },
        "/barber/{id}": {
            "get": {
                "description": "Get barber by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "barber"
                ],
                "summary": "Get barber",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Barber ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/barber.Output"
                        }
                    }
                }
            },
            "put": {
                "description": "Update barber by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "barber"
                ],
                "summary": "Update barber",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Barber ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Barber update input",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/barber.Input"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/barber.Output"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete barber by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "barber"
                ],
                "summary": "Delete barber",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Barber ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    }
                }
            }
        },
        "/barber/{id}/checkin": {
            "get": {
                "description": "Get check-ins for a barber",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "barber"
                ],
                "summary": "Get check-ins",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Barber ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Initial date",
                        "name": "initial_date",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Final date",
                        "name": "final_date",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/barber.PagCheckinOutput"
                        }
                    }
                }
            },
            "post": {
                "description": "Add a check-in for a barber",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "barber"
                ],
                "summary": "Add check-in",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Barber ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Check-in input",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/barber.CheckinInput"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/barber.CheckinOutput"
                        }
                    }
                }
            }
        },
        "/barber/{id}/service": {
            "post": {
                "description": "Add a service to a barber",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "barber"
                ],
                "summary": "Add service to barber",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Barber ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Service input",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/barber.ServiceInput"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    }
                }
            }
        },
        "/barber/{id}/service/{service_id}": {
            "delete": {
                "description": "Remove a service from a barber",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "barber"
                ],
                "summary": "Remove service from barber",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Barber ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Service ID",
                        "name": "service_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    }
                }
            }
        },
        "/service": {
            "get": {
                "description": "List all services",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "service"
                ],
                "summary": "List services",
                "parameters": [
                    {
                        "type": "string",
                        "description": "full or partial service name",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "barber ID",
                        "name": "barber_id",
                        "in": "query"
                    },
                    {
                        "type": "boolean",
                        "description": "is combo",
                        "name": "is_combo",
                        "in": "query"
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "collectionFormat": "csv",
                        "description": "kinds",
                        "name": "kinds",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/service.PagOutput"
                        }
                    }
                }
            },
            "post": {
                "description": "Add a new service",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "service"
                ],
                "summary": "Add service",
                "parameters": [
                    {
                        "description": "Service input",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/service.Input"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/service.Output"
                        }
                    }
                }
            }
        },
        "/service/{id}": {
            "get": {
                "description": "Get service by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "service"
                ],
                "summary": "Get service",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Service ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/service.Output"
                        }
                    }
                }
            },
            "put": {
                "description": "Update service by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "service"
                ],
                "summary": "Update service",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Service ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Service update input",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/service.Input"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/service.Output"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete service by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "service"
                ],
                "summary": "Delete service",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Service ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    }
                }
            }
        }
    },
    "definitions": {
        "admin.Input": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "admin.Output": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "admin.PagOutput": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/admin.Output"
                    }
                },
                "next": {
                    "type": "boolean"
                }
            }
        },
        "admin.UpdateInput": {
            "type": "object",
            "required": [
                "password"
            ],
            "properties": {
                "password": {
                    "type": "string"
                }
            }
        },
        "auth.LoginInput": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "auth.LoginOutput": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "barber.CheckinInput": {
            "type": "object",
            "required": [
                "date_time"
            ],
            "properties": {
                "barberID": {
                    "type": "integer"
                },
                "date_time": {
                    "type": "string"
                }
            }
        },
        "barber.CheckinOutput": {
            "type": "object",
            "properties": {
                "barber_id": {
                    "type": "integer"
                },
                "date_time": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                }
            }
        },
        "barber.Input": {
            "type": "object",
            "required": [
                "name",
                "photo_url"
            ],
            "properties": {
                "commission_rate": {
                    "type": "number"
                },
                "name": {
                    "type": "string"
                },
                "photo_url": {
                    "type": "string"
                }
            }
        },
        "barber.Output": {
            "type": "object",
            "properties": {
                "commission_rate": {
                    "type": "number"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "photo_url": {
                    "type": "string"
                }
            }
        },
        "barber.PagCheckinOutput": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/barber.CheckinOutput"
                    }
                },
                "next": {
                    "type": "boolean"
                }
            }
        },
        "barber.PagOutput": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/barber.Output"
                    }
                },
                "next": {
                    "type": "boolean"
                }
            }
        },
        "barber.ServiceInput": {
            "type": "object",
            "required": [
                "services"
            ],
            "properties": {
                "services": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                }
            }
        },
        "service.Input": {
            "type": "object",
            "required": [
                "duration",
                "kinds",
                "name",
                "price"
            ],
            "properties": {
                "commission_rate": {
                    "type": "number"
                },
                "description": {
                    "type": "string"
                },
                "duration": {
                    "type": "integer"
                },
                "is_combo": {
                    "type": "boolean"
                },
                "kinds": {
                    "type": "array",
                    "items": {
                        "type": "string",
                        "enum": [
                            "haircut",
                            "shave",
                            "beard",
                            "eyebrow"
                        ]
                    }
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                }
            }
        },
        "service.Output": {
            "type": "object",
            "properties": {
                "commission_rate": {
                    "type": "number"
                },
                "description": {
                    "type": "string"
                },
                "duration": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "is_combo": {
                    "type": "boolean"
                },
                "kinds": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                }
            }
        },
        "service.PagOutput": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/service.Output"
                    }
                },
                "next": {
                    "type": "boolean"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:4002",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "My BarberShop API",
	Description:      "API for a My BarberShop application",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
