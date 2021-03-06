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
            "name": "orangeboyChen"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/course": {
            "post": {
                "security": [
                    {
                        "ApiAuthToken": []
                    }
                ],
                "tags": [
                    "菜品"
                ],
                "summary": "上传菜品",
                "parameters": [
                    {
                        "description": "菜品详情",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CourseDto"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/course/query": {
            "get": {
                "security": [
                    {
                        "ApiAuthToken": []
                    }
                ],
                "tags": [
                    "菜品"
                ],
                "summary": "查找菜品",
                "parameters": [
                    {
                        "type": "string",
                        "description": "查找类型，可选: tag",
                        "name": "by",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "关键词",
                        "name": "keyword",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "页数",
                        "name": "pageNum",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "页面大小",
                        "name": "pageSize",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/course/recommend": {
            "get": {
                "security": [
                    {
                        "ApiAuthToken": []
                    }
                ],
                "tags": [
                    "菜品"
                ],
                "summary": "获取推荐列表",
                "responses": {}
            }
        },
        "/course/search": {
            "get": {
                "security": [
                    {
                        "ApiAuthToken": []
                    }
                ],
                "description": "根据关键字搜索菜品",
                "tags": [
                    "菜品"
                ],
                "summary": "搜索菜品",
                "parameters": [
                    {
                        "type": "string",
                        "description": "关键词",
                        "name": "keyword",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "页数",
                        "name": "pageNum",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "页面大小",
                        "name": "pageSize",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Result"
                        }
                    }
                }
            }
        },
        "/course/{courseId}": {
            "get": {
                "security": [
                    {
                        "ApiAuthToken": []
                    }
                ],
                "tags": [
                    "菜品"
                ],
                "summary": "获取菜品详情",
                "parameters": [
                    {
                        "type": "string",
                        "description": "菜品id",
                        "name": "courseId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            },
            "put": {
                "security": [
                    {
                        "ApiAuthToken": []
                    }
                ],
                "tags": [
                    "菜品"
                ],
                "summary": "更新菜品",
                "parameters": [
                    {
                        "type": "string",
                        "description": "courseId",
                        "name": "courseId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "菜品详情",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CourseDto"
                        }
                    }
                ],
                "responses": {}
            },
            "delete": {
                "security": [
                    {
                        "ApiAuthToken": []
                    }
                ],
                "tags": [
                    "菜品"
                ],
                "summary": "删除菜品",
                "parameters": [
                    {
                        "type": "string",
                        "description": "courseId",
                        "name": "courseId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/login": {
            "post": {
                "tags": [
                    "鉴权"
                ],
                "summary": "登录",
                "parameters": [
                    {
                        "description": "登录数据",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UserLoginDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Result"
                        }
                    }
                }
            }
        },
        "/tag/type/list": {
            "get": {
                "security": [
                    {
                        "ApiAuthToken": []
                    }
                ],
                "tags": [
                    "标签"
                ],
                "summary": "获取标签种类列表",
                "responses": {}
            }
        },
        "/tag/type/{tagTypeId}": {
            "get": {
                "security": [
                    {
                        "ApiAuthToken": []
                    }
                ],
                "tags": [
                    "标签"
                ],
                "summary": "获取标签列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "tagTypeId",
                        "name": "typeId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/user": {
            "get": {
                "security": [
                    {
                        "ApiAuthToken": []
                    }
                ],
                "tags": [
                    "用户"
                ],
                "summary": "获取用户信息",
                "responses": {}
            }
        },
        "/user/avatar": {
            "put": {
                "security": [
                    {
                        "ApiAuthToken": []
                    }
                ],
                "tags": [
                    "用户"
                ],
                "summary": "上传头像",
                "parameters": [
                    {
                        "type": "file",
                        "description": "头像",
                        "name": "avatar",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/user/avatar/{avatarFileName}": {
            "get": {
                "security": [
                    {
                        "ApiAuthToken": []
                    }
                ],
                "tags": [
                    "用户"
                ],
                "summary": "获取用户头像",
                "parameters": [
                    {
                        "type": "string",
                        "description": "avatarFileName",
                        "name": "avatarFileName",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/user/info": {
            "put": {
                "security": [
                    {
                        "ApiAuthToken": []
                    }
                ],
                "tags": [
                    "用户"
                ],
                "summary": "更新用户信息",
                "parameters": [
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UserInfoDto"
                        }
                    }
                ],
                "responses": {}
            }
        }
    },
    "definitions": {
        "dto.CourseDto": {
            "type": "object",
            "properties": {
                "detail": {
                    "type": "string"
                },
                "image": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "step": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.CourseStepDto"
                    }
                },
                "tags": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "id1",
                        "id2"
                    ]
                }
            }
        },
        "dto.CourseStepDto": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "order": {
                    "type": "integer"
                },
                "second": {
                    "type": "integer"
                }
            }
        },
        "dto.UserInfoDto": {
            "type": "object",
            "properties": {
                "birthday": {
                    "type": "integer"
                },
                "gender": {
                    "type": "integer"
                },
                "nickName": {
                    "type": "string",
                    "example": "傻逼"
                }
            }
        },
        "dto.UserLoginDto": {
            "type": "object",
            "properties": {
                "identityToken": {
                    "type": "string",
                    "example": "你他妈向苹果登录那个token"
                }
            }
        },
        "response.Result": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "message": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiAuthToken": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "不叫外卖后端API文档",
	Description:      "如有问题，请联系orangeboy",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
