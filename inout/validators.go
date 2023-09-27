package inout

import (
	"errors"
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/indahouse-aleksandr/mp_backend/validators"
)

type OutValidatorsErr struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

const DefaultErrMsg = "unknown error"

var ErrMsgByTag map[string]func(fe validator.FieldError) string = map[string]func(fe validator.FieldError) string{
	"required":             MsgErrHandlerReqired,
	"lte":                  MsgErrHandlerLte,
	"gte":                  MsgErrHandlerGte,
	"email":                MsgErrHandlerEmail,
	validators.TagMyCustom: MsgErrHandlerMyCustom,
}

func ErrMsgValidator(err error) []OutValidatorsErr {
	var out []OutValidatorsErr
	var arrFieldError validator.ValidationErrors

	if !errors.Is(err, arrFieldError) {
		out = append(out, OutValidatorsErr{
			Field:   reflect.TypeOf(err).String(),
			Message: err.Error(),
		})
		return out
	}

	if errors.Is(err, arrFieldError); errors.As(err, &arrFieldError) {
		out := make([]OutValidatorsErr, len(arrFieldError))
		for index, fieldError := range arrFieldError {
			out[index] = OutValidatorsErr{
				Field:   fieldError.Field(),
				Message: getErrorMsg(fieldError),
			}
		}
		return out
	}

	return out
}

func getErrorMsg(fieldError validator.FieldError) string {
	msgHandler, isTagPresent := ErrMsgByTag[fieldError.Tag()]
	if !isTagPresent {
		return DefaultErrMsg
	}

	return msgHandler(fieldError)
}

func MsgErrHandlerReqired(fieldError validator.FieldError) string {
	return "This field is required"
}

func MsgErrHandlerLte(fieldError validator.FieldError) string {
	return "Should be less than " + fieldError.Field()
}

func MsgErrHandlerGte(fieldError validator.FieldError) string {
	return "Should be greater than " + fieldError.Field()
}

func MsgErrHandlerEmail(fieldError validator.FieldError) string {
	return "Email not valid"
}

func MsgErrHandlerMyCustom(fieldError validator.FieldError) string {
	return "My custom error message"
}
