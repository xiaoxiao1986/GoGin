// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-02-07 11:55:03.4541578 -0600 STD m=+0.053236201

package docs

import (
	"bytes"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "swagger": "2.0",
    "info": {
        "description": "Golang Gin API Boilerplate",
        "title": "Gogin API",
        "termsOfService": "https://github.com/xxphenomxx/gogin",
        "contact": {},
        "license": {
            "name": "MIT",
            "url": "https://github.com/xxphenomxx/gogin/blob/master/LICENSE"
        },
        "version": "1.0"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/tags/import": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "Import Image",
                "parameters": [
                    {
                        "type": "file",
                        "description": "Image File",
                        "name": "image",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/app.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/app.Response"
                        }
                    }
                }
            }
        },
        "/user/all": {
            "get": {
                "description": "Returns an array of all User objects",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Get users",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/app.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/app.Response"
                        }
                    }
                }
            }
        },
        "/user/email/{email}": {
            "get": {
                "description": "Queries for a single User object by the email address provided otherwise will return an error",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Fetch a user record by email address",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Email",
                        "name": "email",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/app.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/app.Response"
                        }
                    }
                }
            }
        },
        "/user/get-or-create": {
            "post": {
                "description": "Queries for a single User object by the email address provided, if it cannot locate one then it will create a new User record and return it",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Get or Create a User object - returns a User object and JWT token",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Email",
                        "name": "email",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Google ID",
                        "name": "gID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Image URL",
                        "name": "image",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Users full name",
                        "name": "full",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "First name",
                        "name": "fn",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Last name",
                        "name": "ln",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/app.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/app.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "app.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "type": "object"
                },
                "msg": {
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo swaggerInfo

type s struct{}

func (s *s) ReadDoc() string {
	t, err := template.New("swagger_info").Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, SwaggerInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
