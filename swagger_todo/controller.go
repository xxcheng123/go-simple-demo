package main

import "github.com/gin-gonic/gin"

// RegisterHandler 处理注册
// @Summary 处理用户注册
// @Description 处理用户注册的Controller
// @Tags userRegister
// @Accept application/json
// @Produce application/json
// @Param param body ParamUserCommon true "data"
// @Response 200 {object} ResponseWrapper
// @Router /api/register.do [post]
func RegisterHandler(ctx *gin.Context) {
	param := new(ParamUserCommon)
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ResponseError(ctx, CodeErrorParamError)
		return
	}
	u := new(UserInfo)
	u.Username = param.Username
	ResponseSuccess(ctx, u)
}
