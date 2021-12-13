package helpers

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate


func  ValidateBody(models interface{}) error  {
	validate = validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	err := validate.Struct(models)
	return err
}