package main

import (
	"fmt"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"testing"
)

func Test_Validator(t *testing.T) {
	type User struct {
		UserName string `json:"user_name" validate:"required"`
		Password string `json:"password" validate:"required,min=6,max=20"`
	}
	u := &User{
		Password: "123",
	}
	validate := validator.New()

	if errs := validate.Struct(u); errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			fmt.Println(err)
		}
	}
}
func Test_Validator_Local(t *testing.T) {
	type User struct {
		UserName string `json:"user_name" validate:"required"`
		Password string `json:"password" validate:"required,min=6,max=20"`
	}
	u := &User{
		Password: "123",
	}
	uni := ut.New(zh.New())
	zhTranslator, _ := uni.GetTranslator("zh")
	//if !found {
	//	fmt.Println("zhTranslator not found.", found)
	//	return
	//}
	validate := validator.New()
	if err := zh_translations.RegisterDefaultTranslations(validate, zhTranslator); err != nil {
		fmt.Println(err)
	}
	if errs := validate.Struct(u); errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			fmt.Println(err.Translate(zhTranslator))
		}
	}
}
