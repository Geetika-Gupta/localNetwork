package customvalidator

import (
	"reflect"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v8"
)

//RegisterValidations add the custom validations.
func RegisterValidations() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("devicetype", deviceType)
	}
}

//Validation for device type
func deviceType(v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
	val := strings.ToLower(field.String())
	if val == "computer" || val == "repeater" {
		return true
	}
	return false
}
