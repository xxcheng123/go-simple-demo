package main

import "github.com/gin-gonic/gin"

func ResponseSuccess(ctx *gin.Context, data ResponseData) {
	ctx.JSON(200, &ResponseWrapper{
		Code:    CodeSuccess,
		Message: CodeSuccess.GetMsg(),
		Data:    data,
	})
}

func ResponseError(ctx *gin.Context, code ResponseCode) {
	ctx.JSON(200, &ResponseWrapper{
		Code:    code,
		Message: code.GetMsg(),
		Data:    nil,
	})
}
