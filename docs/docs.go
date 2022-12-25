// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Davydov Mikhail",
            "url": "https://github.com/faringet/City_Temp_Rest_Api",
            "email": "mik.davydov14@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/": {
            "get": {
                "description": "Takes a citys and temperature JSON from DB. Return all city's weather temperature.",
                "summary": "Get all citys from subscription",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Sub"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Takes a city JSON and store in DB. Return city's weather temperature.",
                "summary": "Add a new city to a subscription",
                "parameters": [
                    {
                        "description": "Sub JSON",
                        "name": "city",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Sub"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Sub"
                        }
                    }
                }
            }
        },
        "/{id}": {
            "get": {
                "description": "Takes a city and temperature JSON from DB. Return city weather temperature.",
                "summary": "Get single city from subscription",
                "parameters": [
                    {
                        "type": "string",
                        "description": "search city by id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Sub"
                        }
                    }
                }
            },
            "put": {
                "description": "Update single city.",
                "summary": "Update single city",
                "parameters": [
                    {
                        "type": "string",
                        "description": "search city by id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Sub JSON",
                        "name": "city",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Sub"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Sub"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a city and temperature from DB.",
                "summary": "Delete single city from subscription",
                "parameters": [
                    {
                        "type": "string",
                        "description": "search city by id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Sub"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Sub": {
            "type": "object"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:3000",
	BasePath:         "/subs",
	Schemes:          []string{},
	Title:            "Temperature City Service",
	Description:      "A weather service API in Go using Gin framework.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
