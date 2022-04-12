// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
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
        "/tenant/api/v1/callback": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "获取access_token",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "tenant"
                ],
                "summary": "获取access_token",
                "parameters": [
                    {
                        "type": "string",
                        "description": "code",
                        "name": "code",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "state",
                        "name": "state",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "error",
                        "name": "error",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "error_description",
                        "name": "error_description",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/form.TenantCallbackResp"
                        }
                    }
                }
            }
        },
        "/tenant/api/v1/client": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "获取client配置",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "tenant"
                ],
                "summary": "获取client配置",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/form.TenantClientResp"
                        }
                    }
                }
            }
        },
        "/tenant/api/v1/current-user": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "获取当前用户",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "tenant"
                ],
                "summary": "获取当前用户",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/auth.Claims"
                        }
                    }
                }
            }
        },
        "/tenant/api/v1/refresh-token": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "获取access_token",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "tenant"
                ],
                "summary": "获取access_token",
                "parameters": [
                    {
                        "type": "string",
                        "description": "refresh_token",
                        "name": "refresh_token",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/form.TenantCallbackResp"
                        }
                    }
                }
            }
        },
        "/tenant/api/v1/tenant": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "列表tenant",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "tenant"
                ],
                "summary": "列表tenant",
                "parameters": [
                    {
                        "type": "string",
                        "description": "租户名称",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "当前页",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "每页容量",
                        "name": "pageSize",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Page"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/form.TenantListItem"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "创建tenant",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "tenant"
                ],
                "summary": "创建tenant",
                "parameters": [
                    {
                        "description": "data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/form.TenantCreateReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/tenant/api/v1/tenant/{id}": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "获取tenant",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "tenant"
                ],
                "summary": "获取tenant",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/form.TenantGetResp"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "更新tenant",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "tenant"
                ],
                "summary": "更新tenant",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/form.TenantUpdateReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "删除tenant",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "tenant"
                ],
                "summary": "删除tenant",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "auth.Claims": {
            "type": "object",
            "properties": {
                "accessToken": {
                    "type": "string"
                },
                "address": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "affiliation": {
                    "type": "string"
                },
                "aud": {
                    "description": "the ` + "`" + `aud` + "`" + ` (Audience) claim. See https://datatracker.ietf.org/doc/html/rfc7519#section-4.1.3",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "avatar": {
                    "type": "string"
                },
                "bio": {
                    "type": "string"
                },
                "birthday": {
                    "type": "string"
                },
                "createdIp": {
                    "type": "string"
                },
                "createdTime": {
                    "type": "string"
                },
                "dingtalk": {
                    "type": "string"
                },
                "displayName": {
                    "type": "string"
                },
                "education": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "exp": {
                    "description": "the ` + "`" + `exp` + "`" + ` (Expiration Time) claim. See https://datatracker.ietf.org/doc/html/rfc7519#section-4.1.4",
                    "$ref": "#/definitions/jwt.NumericDate"
                },
                "facebook": {
                    "type": "string"
                },
                "gender": {
                    "type": "string"
                },
                "gitee": {
                    "type": "string"
                },
                "github": {
                    "type": "string"
                },
                "gitlab": {
                    "type": "string"
                },
                "google": {
                    "type": "string"
                },
                "hash": {
                    "type": "string"
                },
                "homepage": {
                    "type": "string"
                },
                "iat": {
                    "description": "the ` + "`" + `iat` + "`" + ` (Issued At) claim. See https://datatracker.ietf.org/doc/html/rfc7519#section-4.1.6",
                    "$ref": "#/definitions/jwt.NumericDate"
                },
                "id": {
                    "type": "string"
                },
                "idCard": {
                    "type": "string"
                },
                "idCardType": {
                    "type": "string"
                },
                "isAdmin": {
                    "type": "boolean"
                },
                "isDefaultAvatar": {
                    "type": "boolean"
                },
                "isDeleted": {
                    "type": "boolean"
                },
                "isForbidden": {
                    "type": "boolean"
                },
                "isGlobalAdmin": {
                    "type": "boolean"
                },
                "isOnline": {
                    "type": "boolean"
                },
                "iss": {
                    "description": "the ` + "`" + `iss` + "`" + ` (Issuer) claim. See https://datatracker.ietf.org/doc/html/rfc7519#section-4.1.1",
                    "type": "string"
                },
                "jti": {
                    "description": "the ` + "`" + `jti` + "`" + ` (JWT ID) claim. See https://datatracker.ietf.org/doc/html/rfc7519#section-4.1.7",
                    "type": "string"
                },
                "karma": {
                    "type": "integer"
                },
                "language": {
                    "type": "string"
                },
                "lark": {
                    "type": "string"
                },
                "lastSigninIp": {
                    "type": "string"
                },
                "lastSigninTime": {
                    "type": "string"
                },
                "ldap": {
                    "type": "string"
                },
                "linkedin": {
                    "type": "string"
                },
                "location": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "nbf": {
                    "description": "the ` + "`" + `nbf` + "`" + ` (Not Before) claim. See https://datatracker.ietf.org/doc/html/rfc7519#section-4.1.5",
                    "$ref": "#/definitions/jwt.NumericDate"
                },
                "owner": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "passwordSalt": {
                    "type": "string"
                },
                "permanentAvatar": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "preHash": {
                    "type": "string"
                },
                "properties": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "qq": {
                    "type": "string"
                },
                "ranking": {
                    "type": "integer"
                },
                "region": {
                    "type": "string"
                },
                "score": {
                    "type": "integer"
                },
                "signupApplication": {
                    "type": "string"
                },
                "sub": {
                    "description": "the ` + "`" + `sub` + "`" + ` (Subject) claim. See https://datatracker.ietf.org/doc/html/rfc7519#section-4.1.2",
                    "type": "string"
                },
                "tag": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                },
                "updatedTime": {
                    "type": "string"
                },
                "wechat": {
                    "type": "string"
                },
                "wecom": {
                    "type": "string"
                },
                "weibo": {
                    "type": "string"
                }
            }
        },
        "form.TenantCallbackResp": {
            "type": "object",
            "properties": {
                "accessToken": {
                    "description": "AccessToken is the token that authorizes and authenticates\nthe requests.",
                    "type": "string"
                },
                "expiry": {
                    "description": "Expiry is the optional expiration time of the access token.\n\nIf zero, TokenSource implementations will reuse the same\ntoken forever and RefreshToken or equivalent\nmechanisms for that TokenSource will not be used.",
                    "type": "string"
                },
                "refreshToken": {
                    "description": "RefreshToken is a token that's used by the application\n(as opposed to the user) to refresh the access token\nif it expires.",
                    "type": "string"
                },
                "tokenType": {
                    "description": "TokenType is the type of token.\nThe Type method returns either this or \"Bearer\", the default.",
                    "type": "string"
                }
            }
        },
        "form.TenantClientResp": {
            "type": "object",
            "properties": {
                "appName": {
                    "type": "string"
                },
                "clientId": {
                    "type": "string"
                },
                "organizationName": {
                    "type": "string"
                },
                "serverUrl": {
                    "type": "string"
                }
            }
        },
        "form.TenantCreateReq": {
            "type": "object",
            "required": [
                "expiredAt",
                "name"
            ],
            "properties": {
                "contact": {
                    "description": "联系方式",
                    "type": "string"
                },
                "description": {
                    "description": "描述",
                    "type": "string"
                },
                "domains": {
                    "description": "域名",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "email": {
                    "description": "邮箱",
                    "type": "string"
                },
                "expiredAt": {
                    "description": "有效期",
                    "type": "string"
                },
                "name": {
                    "description": "名称",
                    "type": "string"
                },
                "status": {
                    "description": "状态",
                    "type": "integer"
                },
                "system": {
                    "description": "系统管理",
                    "type": "boolean"
                }
            }
        },
        "form.TenantGetResp": {
            "type": "object",
            "required": [
                "expiredAt"
            ],
            "properties": {
                "contact": {
                    "description": "联系方式",
                    "type": "string"
                },
                "createdAt": {
                    "description": "创建时间",
                    "type": "string"
                },
                "description": {
                    "description": "描述",
                    "type": "string"
                },
                "domains": {
                    "description": "域名",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "expiredAt": {
                    "description": "有效期",
                    "type": "string"
                },
                "name": {
                    "description": "名称",
                    "type": "string"
                },
                "status": {
                    "description": "状态",
                    "type": "integer"
                },
                "system": {
                    "description": "系统管理",
                    "type": "boolean"
                },
                "updatedAt": {
                    "description": "更新时间",
                    "type": "string"
                }
            }
        },
        "form.TenantListItem": {
            "type": "object",
            "required": [
                "expiredAt"
            ],
            "properties": {
                "contact": {
                    "description": "联系方式",
                    "type": "string"
                },
                "createdAt": {
                    "description": "创建时间",
                    "type": "string"
                },
                "description": {
                    "description": "描述",
                    "type": "string"
                },
                "domains": {
                    "description": "域名",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "expiredAt": {
                    "description": "有效期",
                    "type": "string"
                },
                "id": {
                    "description": "id",
                    "type": "string"
                },
                "name": {
                    "description": "名称",
                    "type": "string"
                },
                "status": {
                    "description": "状态",
                    "type": "integer"
                },
                "system": {
                    "description": "系统管理",
                    "type": "boolean"
                },
                "updatedAt": {
                    "description": "更新时间",
                    "type": "string"
                }
            }
        },
        "form.TenantUpdateReq": {
            "type": "object",
            "required": [
                "expiredAt"
            ],
            "properties": {
                "contact": {
                    "description": "联系方式",
                    "type": "string"
                },
                "description": {
                    "description": "描述",
                    "type": "string"
                },
                "domains": {
                    "description": "域名",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "email": {
                    "description": "邮箱",
                    "type": "string"
                },
                "expiredAt": {
                    "description": "有效期",
                    "type": "string"
                },
                "name": {
                    "description": "名称",
                    "type": "string"
                }
            }
        },
        "jwt.NumericDate": {
            "type": "object",
            "properties": {
                "time.Time": {
                    "type": "string"
                }
            }
        },
        "response.Page": {
            "type": "object",
            "properties": {
                "current": {
                    "type": "integer"
                },
                "pageSize": {
                    "type": "integer"
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "response.Response": {
            "type": "object",
            "properties": {
                "errorCode": {
                    "type": "string"
                },
                "errorMessage": {
                    "type": "string"
                },
                "host": {
                    "type": "string"
                },
                "showType": {
                    "type": "integer"
                },
                "success": {
                    "type": "boolean"
                },
                "traceId": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "Bearer": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
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
	Version:     "0.0.1",
	Host:        "localhost:9094",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "tenant API",
	Description: "tenant接口文档",
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
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
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
	swag.Register("swagger", &s{})
}
