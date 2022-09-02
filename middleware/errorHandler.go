package middleware

import (
	"github.com/SerkanKutlu/orderService/customerror"
	"github.com/labstack/echo/v4"
)

func ErrorHandler(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		defer func() {
			if err := recover(); err != nil {
				switch err.(type) {
				case *customerror.CustomError:
					ce := err.(*customerror.CustomError)
					_ = c.JSON(ce.StatusCode, ce)
				default:
					ce := customerror.NewError(err, 500)
					_ = c.JSON(ce.StatusCode, ce)
				}
			}
		}()
		return next(c)
	}
}
