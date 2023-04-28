package utils

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"reflect"

	"github.com/go-playground/validator/v10"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

// Validator is a custom validator for the Server
type Validator struct {
	validator *validator.Validate
}

// Validate validates the request body and returns a user facing error if validation fails
func (v *Validator) Validate(i interface{}) error {
	if err := v.validator.Struct(i); err != nil {
		fmt.Println("err", err)
		return convertToUserFacingError(err)
	}

	return nil
}

// NewValidator returns a new Validator
func NewValidator(ctx context.Context) *Validator {
	logger := log.FromContext(ctx).WithName("NewValidator")

	logger.Info("Initializing validator")
	v := validator.New()
	registerCustomTagValidators(v)
	return &Validator{validator: v}
}

func registerCustomTagValidators(v *validator.Validate) {
	err := v.RegisterValidation("base64", base64Validator)
	if err != nil {
		panic(err) // can be suppressed if required
	}
}

func base64Validator(fl validator.FieldLevel) bool {
	if fl.Field().Type().Kind() != reflect.String {
		return false
	}

	return isBase64(fl.Field().String())
}

func isBase64(s string) bool {
	if len(s)%4 != 0 {
		return false
	}

	_, err := base64.StdEncoding.DecodeString(s)
	return err == nil
}

// convertToUserFacingError converts the error(s) to a user facing error
func convertToUserFacingError(err error) error {
	switch t := err.(type) {
	case validator.ValidationErrors:
		var validationErr string
		for _, v := range t {
			validationErr += fmt.Sprintf("%s - %s, ", v.Field(), convertTag(v.Tag(), v.Param()))
		}
		validationErr = "Validation failed for the following field(s): " + validationErr[:len(validationErr)-2]
		return New(http.StatusBadRequest, validationErr)
	case *validator.InvalidValidationError:
		return New(http.StatusBadRequest, "Cannot validate request with body: %s", t.Type)
	default:
		return t
	}
}

// convertTag converts the tag to a meaningful error message
func convertTag(tag, param string) string {
	switch tag {
	case "required":
		return "is required"
	case "gte":
		return "must be greater than or equal to " + param
	case "lte":
		return "must be less than or equal to " + param
	case "oneof":
		return "must be one of the following: " + param
	case "excluded_with_all":
		return "must be excluded"
	case "base64":
		return "must be base64 encoded"
	default:
		return tag
	}
}
