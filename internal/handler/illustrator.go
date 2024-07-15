package handler

import (
	"fmt"
	"github.com/AnthonyFVKT/book-illustrator-api/internal/service"
	pb "github.com/AnthonyFVKT/book-illustrator-srv/proto/illustrator"
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
)

type Illustrator struct {
	srv *service.Illustrator
}

func NewIllustrator(srv *service.Illustrator) *Illustrator {
	return &Illustrator{srv: srv}
}

type IllustrateRequest struct {
	Text string `json:"text" validate:"required"`
}

type IllustrateResponse struct {
	Illustrated []*pb.Illustrated `json:"illustrated"`
}

// Illustrate godoc
// @Summary
// @Tags illustrate
// @Description Receives text, split on parts and generate images.
// @Accept json
// @Param IllustrateRequest body IllustrateRequest true "text"
// @Produce json
// @Success 201 {object} IllustrateResponse "When text processed and image generated successfully"
// @Failure 500 {string} echo.HTTPError "Internal Server Error"
// @Router /illustrate [post]
func (il *Illustrator) Illustrate(c echo.Context) error {
	req := new(IllustrateRequest)
	if err := c.Bind(req); err != nil {
		return fmt.Errorf("invalid request: %s", err.Error())
	}

	if err := c.Validate(req); err != nil {
		return fmt.Errorf("invalid request: %s", err.Error())
	}

	res, err := il.srv.Create(c.Request().Context(), req.Text)
	if err != nil {
		slog.Error(err.Error())
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusCreated, IllustrateResponse{Illustrated: res.Illustrated})
}
