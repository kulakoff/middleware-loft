package mw

import (
	"github.com/labstack/echo/v4"
	"log"
)

func RoleCheck(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		header := ctx.Request().Header.Get("User-Role")

		if header == "admin" {
			log.Printf("is admin")
		}

		err := next(ctx)

		if err != nil {
			return err
		}
		return nil
	}
}
