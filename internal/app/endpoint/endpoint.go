package endpoint

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Service interface {
	DaysLeft() int64
}

type Endpoint struct {
	service Service
}

func New(s Service) *Endpoint {
	return &Endpoint{
		service: s,
	}
}

func (endpoint *Endpoint) Status(ctx echo.Context) error {
	days := endpoint.service.DaysLeft()

	s := fmt.Sprintf("Days left: %v", days)

	err := ctx.JSON(http.StatusOK, s)

	if err != nil {
		return err
	}

	return nil
}
