package validators

import (
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	zhTranslations "gopkg.in/go-playground/validator.v9/translations/zh"
)

var Validate *validator.Validate

func SetUp() {
	zhCn := zh.New()
	uni := ut.New(zhCn)
	trans, _ := uni.GetTranslator("zh")
	Validate = validator.New()
	_ = zhTranslations.RegisterDefaultTranslations(Validate, trans)
}
