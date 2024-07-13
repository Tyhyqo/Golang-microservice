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
        "/logout": {
            "post": {
                "description": "Logout the user by clearing the JWT cookie",
                "tags": [
                    "users"
                ],
                "summary": "Logout the user",
                "responses": {
                    "200": {
                        "description": "logged out successfully",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/order_handler": {
            "post": {
                "description": "Create a new order_handler",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "order_handler"
                ],
                "summary": "Create a new order_handler",
                "parameters": [
                    {
                        "description": "Order",
                        "name": "order_handler",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.Order"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/domain.Order"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/order_handler/{id}": {
            "get": {
                "description": "Get an order_handler by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "order_handler"
                ],
                "summary": "Get an order_handler by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Order ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/domain.Order"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "description": "Update an order_handler by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "order_handler"
                ],
                "summary": "Update an order_handler by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Order ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Order",
                        "name": "order_handler",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.Order"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/domain.Order"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/protected": {
            "get": {
                "description": "Access a protected resource with a valid JWT",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "protected"
                ],
                "summary": "Access a protected resource",
                "responses": {
                    "200": {
                        "description": "protected resource",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "unauthorized",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/users/login": {
            "post": {
                "description": "Login a user_handler",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Login a user_handler",
                "parameters": [
                    {
                        "description": "UserWeb",
                        "name": "user_handler",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.UserWeb"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/users/register": {
            "post": {
                "description": "Register a new user_handler",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Register a new user_handler",
                "parameters": [
                    {
                        "description": "UserWeb",
                        "name": "user_handler",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.UserWeb"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/domain.UserWeb"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.Order": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "customer_name": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "status": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "domain.UserWeb": {
            "type": "object",
            "properties": {
                "hash_password": {
                    "type": "string"
                },
                "isCourier": {
                    "type": "boolean"
                },
                "login": {
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
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
