package http

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/ffauzann/mini-e-wallet/internal/constant"
	"github.com/ffauzann/mini-e-wallet/internal/util"
	goValidator "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

var validator *goValidator.Validate = goValidator.New()

// validateStruct validates and return readable error if its not nil.
func validateStruct(s interface{}) (err error) {
	if err = validator.Struct(s); err != nil {
		validationErrors := goValidator.ValidationErrors{}
		if !errors.As(err, &validationErrors) {
			util.Log().Error(err.Error())
			return
		}

		e := validationErrors[0] // Handle only the first-failed validation.

		switch e.Tag() {
		case "required":
			err = fmt.Errorf("%s is required.", e.Field())
		case "lte":
			err = fmt.Errorf("%s must be lower than or equal %s.", e.Field(), e.Param())
		case "gte":
			err = fmt.Errorf("%s must be greater than or equal %s.", e.Field(), e.Param())
		case "min":
			err = fmt.Errorf("%s must be at least %s characters.", e.Field(), e.Param())
		case "max":
			err = fmt.Errorf("%s must be less than %s characters.", e.Field(), e.Param())
		case "uuid":
			err = fmt.Errorf("%s: malformed uuid.", e.Field())
		}
	}

	return
}

func bindAndValidate(c echo.Context, i interface{}) (code int, err error) {
	if err = c.Bind(i); err != nil {
		util.Log().Error(err.Error())
		return http.StatusUnprocessableEntity, constant.ErrUnprocessableEntity
	}

	if err = validateStruct(i); err != nil {
		util.Log().Error(err.Error())
		return http.StatusUnprocessableEntity, err
	}

	return
}
