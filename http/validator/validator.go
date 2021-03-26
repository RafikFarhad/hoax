package validator

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

// singleton
//var globalValidator *validator.Validate
//var translator *ut.Translator
//var validatorInit sync.Once
//var translatorInit sync.Once
//
//func GetValidator() *validator.Validate {
//	validatorInit.Do(func() {
//		globalValidator = validator.New()
//	})
//	return globalValidator
//}
//
//func GetTranslator() *ut.Translator {
//	translatorInit.Do(func() {
//		english := en.New()
//		universalTranslator := ut.New(english, english)
//		trans, _ := universalTranslator.GetTranslator("en")
//		translator = &trans
//		_ = en_translations.RegisterDefaultTranslations(globalValidator, trans)
//	})
//	return translator
//}

func GetValidator() *validator.Validate {
	return validator.New()
}

func GetTranslator() *ut.Translator {
	english := en.New()
	universalTranslator := ut.New(english, english)
	trans, _ := universalTranslator.GetTranslator("en")
	translator := &trans
	_ = en_translations.RegisterDefaultTranslations(GetValidator(), trans)
	return translator
}
