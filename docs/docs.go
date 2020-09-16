// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

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
        "contact": {},
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/adminUser": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "账户管理"
                ],
                "summary": "获取账户列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "姓名|手机",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "页码",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "每页显示多少条",
                        "name": "pageSize",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.AdminUserPage"
                        }
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "账户管理"
                ],
                "summary": "添加账户",
                "parameters": [
                    {
                        "description": "用户",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.UserAdd"
                        }
                    }
                ]
            }
        },
        "/adminUser/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "账户管理"
                ],
                "summary": "获取账户详情",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "用户ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.AdminUserList"
                        }
                    }
                }
            },
            "put": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "账户管理"
                ],
                "summary": "编辑账户",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "用户ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "用户",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.UserEdit"
                        }
                    }
                ]
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "账户管理"
                ],
                "summary": "删除账户",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "用户ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ]
            }
        },
        "/auth/role": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "账户管理"
                ],
                "summary": "获取角色树",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.AdminUserPage"
                        }
                    }
                }
            }
        },
        "/auth/tree": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "角色管理"
                ],
                "summary": "获取权限树",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.RolePage"
                        }
                    }
                }
            }
        },
        "/captcha": {
            "get": {
                "tags": [
                    "登录操作"
                ],
                "summary": "验证码",
                "responses": {
                    "200": {}
                }
            }
        },
        "/index": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "管理员首页"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.SystemMessagePage"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "登录操作"
                ],
                "summary": "登录操作",
                "parameters": [
                    {
                        "description": "用户",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.LoginUser"
                        }
                    }
                ],
                "responses": {
                    "200": {}
                }
            }
        },
        "/role": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "角色管理"
                ],
                "summary": "获取所有角色",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "页码",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "每页显示多少条",
                        "name": "pageSize",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.RolePage"
                        }
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "角色管理"
                ],
                "summary": "添加角色",
                "parameters": [
                    {
                        "description": "角色",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.RoleAdd"
                        }
                    }
                ],
                "responses": {
                    "200": {}
                }
            }
        },
        "/role/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "角色管理"
                ],
                "summary": "获取角色详情",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "角色ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.RoleList"
                        }
                    }
                }
            },
            "put": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "角色管理"
                ],
                "summary": "编辑角色",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "角色ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "角色",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.RoleEdit"
                        }
                    }
                ]
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "角色管理"
                ],
                "summary": "删除角色",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "角色ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {}
                }
            }
        },
        "/systemMessage": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "系统公告"
                ],
                "summary": "获取系统公告列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "标题",
                        "name": "title",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "页码",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "每页显示多少条",
                        "name": "pageSize",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.SystemMessagePage"
                        }
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "系统公告"
                ],
                "summary": "添加系统公告",
                "parameters": [
                    {
                        "description": "系统公告",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.MessageAdd"
                        }
                    }
                ],
                "responses": {
                    "200": {}
                }
            }
        },
        "/systemMessage/{id}": {
            "put": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "系统公告"
                ],
                "summary": "编辑公告",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "公告ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "公告",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.MessageEdit"
                        }
                    }
                ]
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "系统公告"
                ],
                "summary": "删除公告",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "公告ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ]
            }
        },
        "/systemMessageByUser": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "系统公告"
                ],
                "summary": "获取用户公告列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "username",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "页码",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "每页显示多少条",
                        "name": "pageSize",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.AdminUserMessagePage"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "request.LoginUser": {
            "type": "object",
            "required": [
                "password",
                "username",
                "verify_code"
            ],
            "properties": {
                "password": {
                    "description": "密码",
                    "type": "string",
                    "example": "admin"
                },
                "username": {
                    "description": "用户名",
                    "type": "string",
                    "example": "admin"
                },
                "verify_code": {
                    "description": "验证码",
                    "type": "string",
                    "example": "9527"
                }
            }
        },
        "request.MessageAdd": {
            "type": "object",
            "required": [
                "beginTime",
                "endTime",
                "title",
                "users"
            ],
            "properties": {
                "beginTime": {
                    "description": "开始时间",
                    "type": "string",
                    "example": "2020-09-15T14:41:46+08:00"
                },
                "endTime": {
                    "description": "结束时间",
                    "type": "string",
                    "example": "2020-09-23T14:41:50+08:00"
                },
                "title": {
                    "description": "系统公告",
                    "type": "string",
                    "example": "这是一个系统公告"
                },
                "users": {
                    "description": "发给用户id",
                    "type": "array",
                    "items": {
                        "type": "integer"
                    },
                    "example": [
                        1,
                        2,
                        3,
                        4,
                        5,
                        6,
                        7
                    ]
                }
            }
        },
        "request.MessageEdit": {
            "type": "object",
            "required": [
                "beginTime",
                "endTime",
                "title",
                "users"
            ],
            "properties": {
                "beginTime": {
                    "description": "开始时间",
                    "type": "string",
                    "example": "2020-09-15T14:41:46+08:00"
                },
                "endTime": {
                    "description": "结束时间",
                    "type": "string",
                    "example": "2020-09-23T14:41:50+08:00"
                },
                "title": {
                    "description": "系统公告",
                    "type": "string",
                    "example": "这是一个系统公告"
                },
                "users": {
                    "description": "发给用户id",
                    "type": "array",
                    "items": {
                        "type": "integer"
                    },
                    "example": [
                        1,
                        2,
                        3,
                        4,
                        5,
                        6,
                        7
                    ]
                }
            }
        },
        "request.RoleAdd": {
            "type": "object",
            "required": [
                "auth",
                "name"
            ],
            "properties": {
                "auth": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    },
                    "example": [
                        1,
                        2,
                        3,
                        4,
                        5,
                        6,
                        7,
                        8,
                        9,
                        10,
                        11,
                        12,
                        13,
                        14,
                        15,
                        16,
                        17,
                        18,
                        19
                    ]
                },
                "name": {
                    "description": "角色名",
                    "type": "string",
                    "example": "程序员"
                },
                "pid": {
                    "description": "上级ID",
                    "type": "integer",
                    "example": 1
                }
            }
        },
        "request.RoleEdit": {
            "type": "object",
            "required": [
                "auth",
                "name"
            ],
            "properties": {
                "auth": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    },
                    "example": [
                        1,
                        2,
                        3,
                        4,
                        5,
                        6,
                        7,
                        8,
                        9,
                        10,
                        11,
                        12,
                        13,
                        14,
                        15,
                        16,
                        17,
                        18,
                        19
                    ]
                },
                "name": {
                    "description": "角色名",
                    "type": "string",
                    "example": "程序员"
                },
                "pid": {
                    "description": "上级ID",
                    "type": "integer",
                    "example": 1
                }
            }
        },
        "request.UserAdd": {
            "type": "object",
            "required": [
                "password",
                "real_name",
                "roles",
                "tel",
                "user_name"
            ],
            "properties": {
                "password": {
                    "description": "密码",
                    "type": "string",
                    "example": "test"
                },
                "real_name": {
                    "description": "真实姓名",
                    "type": "string",
                    "example": "test"
                },
                "roles": {
                    "description": "所属角色",
                    "type": "array",
                    "items": {
                        "type": "integer"
                    },
                    "example": [
                        1,
                        2
                    ]
                },
                "status": {
                    "description": "状态",
                    "type": "integer",
                    "example": 1
                },
                "tel": {
                    "description": "电话号码",
                    "type": "string",
                    "example": "17585534067"
                },
                "user_name": {
                    "description": "账号",
                    "type": "string",
                    "example": "test"
                }
            }
        },
        "request.UserEdit": {
            "type": "object",
            "required": [
                "real_name",
                "roles",
                "tel",
                "user_name"
            ],
            "properties": {
                "password": {
                    "description": "密码（非必填）",
                    "type": "string",
                    "example": "test"
                },
                "real_name": {
                    "description": "真实姓名",
                    "type": "string",
                    "example": "test"
                },
                "roles": {
                    "description": "所属角色",
                    "type": "array",
                    "items": {
                        "type": "integer"
                    },
                    "example": [
                        1,
                        2
                    ]
                },
                "status": {
                    "description": "状态",
                    "type": "integer",
                    "example": 1
                },
                "tel": {
                    "description": "电话号码",
                    "type": "string",
                    "example": "17585534067"
                },
                "user_name": {
                    "description": "账号",
                    "type": "string",
                    "example": "test"
                }
            }
        },
        "response.AdminSystemMessage": {
            "type": "object",
            "properties": {
                "begin_time": {
                    "type": "string"
                },
                "end_time": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "response.AdminUserList": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "real_name": {
                    "description": "真实姓名",
                    "type": "string"
                },
                "roles": {
                    "description": "角色信息",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/response.CasRole"
                    }
                },
                "status": {
                    "description": "用户状态",
                    "type": "integer"
                },
                "tel": {
                    "description": "手机号码",
                    "type": "string"
                },
                "user_name": {
                    "description": "登录名",
                    "type": "string"
                }
            }
        },
        "response.AdminUserMessageList": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "systemMessages": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/response.AdminSystemMessage"
                    }
                }
            }
        },
        "response.AdminUserMessagePage": {
            "type": "object",
            "properties": {
                "current_page": {
                    "description": "每页显示多少条",
                    "type": "integer"
                },
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/response.AdminUserMessageList"
                    }
                },
                "per_page": {
                    "description": "当前页码",
                    "type": "integer"
                },
                "total": {
                    "description": "总共多少页",
                    "type": "integer"
                }
            }
        },
        "response.AdminUserPage": {
            "type": "object",
            "properties": {
                "current_page": {
                    "description": "每页显示多少条",
                    "type": "integer"
                },
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/response.AdminUserList"
                    }
                },
                "per_page": {
                    "description": "当前页码",
                    "type": "integer"
                },
                "total": {
                    "description": "总共多少页",
                    "type": "integer"
                }
            }
        },
        "response.CasRole": {
            "type": "object",
            "properties": {
                "id": {
                    "description": "角色ID",
                    "type": "integer"
                },
                "name": {
                    "description": "角色名",
                    "type": "string"
                }
            }
        },
        "response.RoleList": {
            "type": "object",
            "properties": {
                "auth": {
                    "description": "所有权限arr",
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "base_auth": {
                    "description": "api权限",
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "description": "角色名",
                    "type": "string"
                },
                "parent_name": {
                    "description": "上级角色名",
                    "type": "string"
                },
                "pid": {
                    "description": "上级ID",
                    "type": "integer"
                }
            }
        },
        "response.RolePage": {
            "type": "object",
            "properties": {
                "current_page": {
                    "description": "每页显示多少条",
                    "type": "integer"
                },
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/response.RoleList"
                    }
                },
                "per_page": {
                    "description": "当前页码",
                    "type": "integer"
                },
                "total": {
                    "description": "总共多少页",
                    "type": "integer"
                }
            }
        },
        "response.SystemMessageList": {
            "type": "object",
            "properties": {
                "begin_time": {
                    "description": "开始时间",
                    "type": "string"
                },
                "end_time": {
                    "description": "结束时间",
                    "type": "string"
                },
                "id": {
                    "description": "id",
                    "type": "integer"
                },
                "status": {
                    "description": "是否作废 0：否，1：是",
                    "type": "integer"
                },
                "title": {
                    "description": "标题",
                    "type": "string"
                },
                "users": {
                    "description": "用户名称",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/response.UserNames"
                    }
                }
            }
        },
        "response.SystemMessagePage": {
            "type": "object",
            "properties": {
                "current_page": {
                    "description": "每页显示多少条",
                    "type": "integer"
                },
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/response.SystemMessageList"
                    }
                },
                "per_page": {
                    "description": "当前页码",
                    "type": "integer"
                },
                "total": {
                    "description": "总共多少页",
                    "type": "integer"
                }
            }
        },
        "response.UserNames": {
            "type": "object",
            "properties": {
                "id": {
                    "description": "用户ID",
                    "type": "integer"
                },
                "user_name": {
                    "description": "用户名",
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
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "y2pay",
	Description: "",
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
