package main

import (
	"github.com/gin-gonic/gin"
	_ "swagger_todo/docs" // main 文件中导入 docs 包

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// gin-swagger middleware

// @title 用户社区模块
// @version 1.0.0
// @description 响应用户注册、登录等

// @license.name Apache 2.0

// @host 127.0.0.1:9998

func main() {
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/api/register.do", RegisterHandler)
	r.Run(":9998")
}
