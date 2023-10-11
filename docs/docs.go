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
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "Health and Readiness"
                ],
                "summary": "Health check endpoint for k8s",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/init": {
            "post": {
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account Management"
                ],
                "summary": "Create new account",
                "parameters": [
                    {
                        "type": "string",
                        "default": "49437636-fa79-40fb-b5cf-5f066235fdda",
                        "description": "uuid",
                        "name": "customer_xid",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/readiness": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "Health and Readiness"
                ],
                "summary": "Readiness endpoint for k8s",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/wallet": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account Management"
                ],
                "summary": "Get balance",
                "parameters": [
                    {
                        "type": "string",
                        "default": "Token xxx",
                        "description": "Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {}
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account Management"
                ],
                "summary": "Enable wallet",
                "parameters": [
                    {
                        "type": "string",
                        "default": "Token xxx",
                        "description": "Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {}
            },
            "patch": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account Management"
                ],
                "summary": "Disable wallet",
                "parameters": [
                    {
                        "type": "string",
                        "default": "Token xxx",
                        "description": "Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/wallet/deposits": {
            "post": {
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Transaction Management"
                ],
                "summary": "Make a deposit",
                "parameters": [
                    {
                        "type": "string",
                        "default": "Token xxx",
                        "description": "Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "49437636-fa79-40fb-b5cf-5f066235fdda",
                        "description": "UUID",
                        "name": "reference_id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "number",
                        "default": 10000,
                        "description": "Deposit Amount",
                        "name": "amount",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/wallet/transactions": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Transaction Management"
                ],
                "summary": "Get transaction history",
                "parameters": [
                    {
                        "type": "string",
                        "default": "Token xxx",
                        "description": "Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/wallet/withdrawals": {
            "post": {
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Transaction Management"
                ],
                "summary": "Request a withdrawal",
                "parameters": [
                    {
                        "type": "string",
                        "default": "Token xxx",
                        "description": "Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "49437636-fa79-40fb-b5cf-5f066235fdda",
                        "description": "UUID",
                        "name": "reference_id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "number",
                        "default": 10000,
                        "description": "Withdrawal Amount",
                        "name": "amount",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {}
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:2201",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Mini E-Wallet API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
