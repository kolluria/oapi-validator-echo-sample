/*
Copyright 2022 VMware, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package echo

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	errorutils "github.com/oapi-validator-echo-sample/utils/errors"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

// Validator is a custom validator for the Server
type Validator struct {
	validator *validator.Validate
}

// Validate validates the request body and returns a user facing error if validation fails
func (v *Validator) Validate(i interface{}) error {
	if err := v.validator.Struct(i); err != nil {
		return convertToUserFacingError(err)
	}

	return nil
}

// NewValidator returns a new Validator
func NewValidator(ctx context.Context) *Validator {
	logger := log.FromContext(ctx).WithName("NewValidator")

	logger.Info("Initializing validator")
	v := validator.New()
	return &Validator{validator: v}
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
		// Example: Validation failed for the following field(s): DeploymentFailureCount - must be excluded, DeploymentSuccessCount - must be excluded, Testbeds - must be excluded
		return errorutils.New(http.StatusBadRequest, validationErr)
	case *validator.InvalidValidationError:
		return errorutils.New(http.StatusBadRequest, "Cannot validate request with body: %s", t.Type)
	default:
		return t
	}
}

// convertTag converts the tag to a meaningful error message
func convertTag(tag, param string) string {
	switch tag {
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
