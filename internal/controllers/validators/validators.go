package validators

import (
	"fmt"
	"reflect"
	"strings"

	enLoc "github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/translations/en"
	"github.com/rs/zerolog"
)

// Validators and Translator
type Validators struct {
	Validate *validator.Validate
	Trans    ut.Translator
	Logger   *zerolog.Logger
}

// New instantiates struct Validate and a message Translators
func New(logger *zerolog.Logger) Validators {
	v := Validators{
		Validate: newTagNameValidator(),
		Trans:    newEnglishTranslator(),
		Logger:   logger,
	}
	RegisterDefaultTranslations(v, logger)

	return v
}

// RegisterDefaultTranslations from built-in library
func RegisterDefaultTranslations(v Validators, logger *zerolog.Logger) {
	if err := en.RegisterDefaultTranslations(v.Validate, v.Trans); err != nil {
		logger.Panic().AnErr("failed to RegisterDefaultTranslations default translations", err)
	}
}

// RegisterStructValidation just wraps Validate for code tidiness
func (v *Validators) RegisterStructValidation(fn validator.StructLevelFunc, types ...interface{}) {
	v.Validate.RegisterStructValidation(fn, types...)
}

// AddTranslation to be used when reporting validation errors
func (v *Validators) AddTranslation(tag string, errMessage string) {
	registerFn := func(ut ut.Translator) error {
		return ut.Add(tag, errMessage, false)
	}
	transFn := func(ut ut.Translator, fe validator.FieldError) string {
		param := fe.Param()
		if param == "" {
			switch fe.Value().(type) {
			case string:
				param = fe.Value().(string)
			default:
				param = "n/a?"
			}
		}
		fieldErrorTag := fe.Tag()
		field := fe.Field()
		t, err := ut.T(fieldErrorTag, field, param)
		if err != nil {
			return fe.(error).Error()
		}
		return t
	}
	_ = v.Validate.RegisterTranslation(tag, v.Trans, registerFn, transFn)
}

// Translate validation errors from machine to human-readable form
func (v *Validators) Translate(err error) error {
	if err == nil {
		return fmt.Errorf("nil error when translating validation failure")
	}
	validatorErrs, ok := err.(validator.ValidationErrors)
	if !ok {
		return err
	}
	errs := make([]string, 0, len(validatorErrs))
	for _, e := range validatorErrs {
		errs = append(errs, e.Translate(v.Trans))
	}
	return fmt.Errorf(strings.Join(errs, `; `))
}

func newTagNameValidator() *validator.Validate {
	tagValidator := validator.New()
	tagValidator.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	return tagValidator
}

func newEnglishTranslator() ut.Translator {
	english := enLoc.New()
	uni := ut.New(english, english)
	trans, _ := uni.GetTranslator("en")
	return trans
}
