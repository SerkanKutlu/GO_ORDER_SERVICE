package middleware

import (
	"github.com/SerkanKutlu/orderService/customerror"
	"github.com/labstack/echo/v4"
)

func ErrorHandler(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		defer func() {
			if err := recover(); err != nil {
				customError := customerror.NewError(err, 500)
				_ = c.JSON(customError.StatusCode, customError)
			}
		}()
		return next(c)
	}
}
