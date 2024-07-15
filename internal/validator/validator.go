package validator

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
)

type StructValidator struct {
	validator *validator.Validate
}

func NewStructValidator() (*StructValidator, error) {
	v := validator.New()
	return &StructValidator{validator: v}, nil
}

func (cv *StructValidator) Validate(i any) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}
