package validatator

import (
	"fmt"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
)

var (
	validate *validator.Validate
	uni      *ut.UniversalTranslator
	trans    ut.Translator
)

func getTrans() ut.Translator {
	if trans == nil {
		en := en.New()
		uni = ut.New(en, en)
		trans, _ = uni.GetTranslator("en")

		enTranslations.RegisterDefaultTranslations(validate, trans)
	}

	return trans
}

func CreateValidationErrorMessage(err error) error {
	for _, e := range err.(validator.ValidationErrors) {
		translatedErr := fmt.Errorf(e.Translate(getTrans()))
		return translatedErr
	}
	return err
}

func GetValidator() *validator.Validate {
	if validate == nil {
		validate = validator.New()
	}
	return validate
}

func Validation(reg interface{}) (err error, check bool) {

	err = GetValidator().Struct(reg)
	if err != nil {
		return CreateValidationErrorMessage(err), true
	}
	return
}
