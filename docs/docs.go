// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/karyawan": {
            "get": {
                "description": "Mengambil seluruh data karyawan",
                "produces": [
                    "application/json"
                ],
                "summary": "Get All Karyawan",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Karyawan"
                            }
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
            "post": {
                "description": "Membuat data karyawan baru",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create Karyawan",
                "parameters": [
                    {
                        "description": "Data karyawan yang ingin dibuat",
                        "name": "karyawan",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Karyawan"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Karyawan"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/karyawan/{id}": {
            "get": {
                "description": "Mengambil data karyawan berdasarkan ID",
                "produces": [
                    "application/json"
                ],
                "summary": "Get Karyawan by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID Karyawan",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Karyawan"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "description": "Mengubah data karyawan berdasarkan ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Update Karyawan",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID Karyawan",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Data karyawan yang ingin diubah",
                        "name": "karyawan",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Karyawan"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Karyawan"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Menghapus data karyawan berdasarkan ID",
                "produces": [
                    "application/json"
                ],
                "summary": "Delete Karyawan",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID Karyawan",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Karyawan": {
            "type": "object"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Data Karyawan API",
	Description:      "API untuk CRUD data karyawan",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}