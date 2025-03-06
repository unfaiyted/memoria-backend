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
        "/health": {
            "get": {
                "description": "returns JSON object with health statuses.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "health"
                ],
                "summary": "checks app and database health",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.HealthResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/paste": {
            "put": {
                "description": "Update a pastes value",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "pastes"
                ],
                "summary": "Update paste",
                "parameters": [
                    {
                        "description": "Updated paste data",
                        "name": "paste",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UpdatePasteRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success response with paste data",
                        "schema": {
                            "$ref": "#/definitions/models.APIResponse-models_PasteData"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Paste not found",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Creates a new paste",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "pastes"
                ],
                "summary": "Create paste",
                "parameters": [
                    {
                        "description": "Paste data",
                        "name": "paste",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreatePasteRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success response with paste data",
                        "schema": {
                            "$ref": "#/definitions/models.APIResponse-models_PasteData"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "delete a paste by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "pastes"
                ],
                "summary": "Deletes paste by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Paste ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.APIResponse-uint64"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/paste/all": {
            "get": {
                "description": "Returns a list of all the pastes. # TODO: pagination",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "pastes"
                ],
                "summary": "Lists out all the pastes",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 1,
                        "description": "Page number",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 10,
                        "description": "Items per page",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success response with paste list data",
                        "schema": {
                            "$ref": "#/definitions/models.APIResponse-models_PasteListData"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/paste/{id}": {
            "get": {
                "description": "Retrieve a paste by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "pastes"
                ],
                "summary": "Gets a specific paste",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Paste ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success response with paste data",
                        "schema": {
                            "$ref": "#/definitions/models.APIResponse-models_PasteData"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Paste not found",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/users": {
            "get": {
                "description": "Get all users in the system",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "List users",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.UserResponse"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new user in the system",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Create a new user",
                "parameters": [
                    {
                        "description": "User data",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.UserResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/users/{id}": {
            "get": {
                "description": "Get a user by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get a user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.UserResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "Update a user's information",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Update a user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "User data",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.UserResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a user by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Delete a user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.APIResponse-models_PasteData": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/models.PasteData"
                },
                "message": {
                    "type": "string",
                    "example": "Operation successful"
                },
                "success": {
                    "type": "boolean",
                    "example": true
                }
            }
        },
        "models.APIResponse-models_PasteListData": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/models.PasteListData"
                },
                "message": {
                    "type": "string",
                    "example": "Operation successful"
                },
                "success": {
                    "type": "boolean",
                    "example": true
                }
            }
        },
        "models.APIResponse-uint64": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "integer"
                },
                "message": {
                    "type": "string",
                    "example": "Operation successful"
                },
                "success": {
                    "type": "boolean",
                    "example": true
                }
            }
        },
        "models.CreatePasteRequest": {
            "type": "object",
            "required": [
                "content",
                "privacy",
                "title"
            ],
            "properties": {
                "content": {
                    "type": "string"
                },
                "expires_at": {
                    "type": "string",
                    "example": "2023-01-08T00:00:00Z"
                },
                "privacy": {
                    "type": "string",
                    "enum": [
                        "public",
                        "private",
                        "password"
                    ]
                },
                "syntax_highlight": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "models.ErrorResponse": {
            "type": "object",
            "properties": {
                "details": {
                    "type": "object",
                    "additionalProperties": true
                },
                "error": {
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.ErrorType"
                        }
                    ],
                    "example": "FAILED_CHECK"
                },
                "message": {
                    "type": "string",
                    "example": "This is a pretty message"
                },
                "request_id": {
                    "type": "string"
                },
                "timestamp": {
                    "type": "string"
                }
            }
        },
        "models.ErrorType": {
            "type": "string",
            "enum": [
                "FAILED_CHECK",
                "UNAUTHORIZED",
                "NOT_FOUND",
                "BAD_REQUEST",
                "INTERNAL_ERROR",
                "FORBIDDEN",
                "CONFLICT",
                "VALIDATION_ERROR",
                "RATE_LIMITED",
                "TIMEOUT",
                "SERVICE_UNAVAILABLE",
                "UNPROCESSABLE_ENTITY"
            ],
            "x-enum-varnames": [
                "ErrorTypeFailedCheck",
                "ErrorTypeUnauthorized",
                "ErrorTypeNotFound",
                "ErrorTypeBadRequest",
                "ErrorTypeInternalError",
                "ErrorTypeForbidden",
                "ErrorTypeConflict",
                "ErrorTypeValidation",
                "ErrorTypeRateLimited",
                "ErrorTypeTimeout",
                "ErrorTypeServiceUnavailable",
                "ErrorTypeUnprocessableEntity"
            ]
        },
        "models.HealthResponse": {
            "type": "object",
            "required": [
                "application",
                "database",
                "status"
            ],
            "properties": {
                "application": {
                    "type": "boolean",
                    "example": true
                },
                "database": {
                    "type": "boolean",
                    "example": true
                },
                "status": {
                    "description": "Overall status of the system",
                    "type": "string",
                    "enum": [
                        "up",
                        "down",
                        "degraded"
                    ],
                    "example": "up"
                }
            }
        },
        "models.Paste": {
            "description": "A text snippet with formatting, expiration, and privacy settings",
            "type": "object",
            "required": [
                "content",
                "id",
                "privacy",
                "syntax_highlight",
                "title"
            ],
            "properties": {
                "content": {
                    "type": "string",
                    "example": "console.log('Hello world');"
                },
                "created_at": {
                    "type": "string",
                    "example": "2023-01-01T00:00:00Z"
                },
                "expires_at": {
                    "type": "string",
                    "example": "2023-01-08T00:00:00Z"
                },
                "id": {
                    "type": "integer",
                    "example": 123111
                },
                "privacy": {
                    "description": "\"public\", \"private\", \"password\"",
                    "type": "string",
                    "enum": [
                        "public",
                        "private",
                        "password"
                    ],
                    "example": "public"
                },
                "syntax_highlight": {
                    "type": "string",
                    "example": "javascript"
                },
                "title": {
                    "type": "string",
                    "example": "My Code Snippet"
                },
                "user_id": {
                    "type": "string",
                    "example": "u98765zyxwv"
                }
            }
        },
        "models.PasteData": {
            "type": "object",
            "properties": {
                "paste": {
                    "$ref": "#/definitions/models.Paste"
                }
            }
        },
        "models.PasteListData": {
            "description": "List of pastes response wrapper",
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "pastes": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Paste"
                    }
                }
            }
        },
        "models.UpdatePasteRequest": {
            "type": "object",
            "required": [
                "content",
                "id",
                "privacy",
                "title"
            ],
            "properties": {
                "content": {
                    "type": "string"
                },
                "expires_at": {
                    "type": "string",
                    "example": "2023-01-08T00:00:00Z"
                },
                "id": {
                    "type": "integer"
                },
                "privacy": {
                    "type": "string",
                    "enum": [
                        "public",
                        "private",
                        "password"
                    ]
                },
                "syntax_highlight": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "models.User": {
            "type": "object",
            "required": [
                "email",
                "name",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "example": "john@example.com"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "name": {
                    "type": "string",
                    "maxLength": 100,
                    "minLength": 2,
                    "example": "John Doe"
                },
                "password": {
                    "description": "omitempty will exclude it from JSON responses",
                    "type": "string",
                    "format": "password",
                    "minLength": 8,
                    "example": "strongpassword123"
                }
            }
        },
        "models.UserResponse": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
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
