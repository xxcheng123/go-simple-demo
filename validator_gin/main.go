package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"strings"
)

type User struct {
	Username string `json:"username" binding:"required,gte=3,lte=8"`
	Password string `json:"password" binding:"required,gte=8,lte=16"`
	Sex      uint8  `json:"sex" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Age      int    `json:"age"`
}

var trans ut.Translator

func main() {
	r := gin.Default()
	if err := InitTrans(); err != nil {
		fmt.Println(err)
		return
	}
	r.POST("/signup.do", func(ctx *gin.Context) {
		u := new(User)
		if err := ctx.ShouldBindJSON(u); err != nil {
			errs, ok := err.(validator.ValidationErrors)
			var msg string = "param is error"
			if ok {
				msg = ""
				for _, validateErr := range errs {
					msg = fmt.Sprintf("%s,%s", msg, validateErr.Translate(trans))
				}
			}
			ctx.JSON(200, msg)
		} else {
			ctx.JSON(200, "signup success")
		}
	})

	r.Run(":9999")
}

func InitTrans() error {
	if validate, ok := binding.Validator.Engine().(*validator.Validate); ok {
		uni := ut.New(zh.New())
		trans, _ = uni.GetTranslator("zh")
		zh_translations.RegisterDefaultTranslations(validate, trans)

		validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})

		return nil
	} else {
		return errors.New("get gin Validator failed")
	}
}
