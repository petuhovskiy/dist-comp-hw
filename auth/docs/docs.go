// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2020-03-29 16:21:48.282872 +0300 MSK m=+0.072192281

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Arthur Petukhovsky",
            "url": "https://t.me/petuhovskiy",
            "email": "petuhovskiy@yandex.ru"
        },
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1/confirm": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Confirms user account phone or email.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Confirmation string",
                        "name": "v",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/modelapi.ConfirmResponse"
                        }
                    }
                }
            }
        },
        "/v1/refresh": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Accepts valid refresh token, returns new refresh and access tokens.",
                "parameters": [
                    {
                        "description": "Refresh token",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/modelapi.RefreshRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/modelapi.IssuedTokens"
                        }
                    }
                }
            }
        },
        "/v1/signin": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Sign in (login)",
                "parameters": [
                    {
                        "description": "Credentials",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/modelapi.SigninRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/modelapi.IssuedTokens"
                        }
                    }
                }
            }
        },
        "/v1/signup": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Sign up (register)",
                "parameters": [
                    {
                        "description": "Credentials",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/modelapi.SignupRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/modelapi.SignupResponse"
                        }
                    }
                }
            }
        },
        "/v1/validate": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Validates access token.",
                "parameters": [
                    {
                        "description": "Access token",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/modelapi.ValidateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/modelapi.ValidateResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "modelapi.ConfirmResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "subject": {
                    "type": "string"
                }
            }
        },
        "modelapi.IssuedTokens": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "access_token_ttl": {
                    "type": "string"
                },
                "refresh_token": {
                    "type": "string"
                },
                "refresh_token_ttl": {
                    "type": "string"
                }
            }
        },
        "modelapi.RefreshRequest": {
            "type": "object",
            "properties": {
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "modelapi.SigninRequest": {
            "type": "object",
            "properties": {
                "login": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "modelapi.SignupRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                }
            }
        },
        "modelapi.SignupResponse": {
            "type": "object",
            "properties": {
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "modelapi.ValidateRequest": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                }
            }
        },
        "modelapi.ValidateResponse": {
            "type": "object",
            "properties": {
                "expire_after": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "localhost:8081",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "Auth API",
	Description: "This is sample auth server, made as dist-comp homework.",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
