package order_handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"tracking_service/internal/domain"
)

// CreateOrder godoc
// @Summary Create a new order_handler
// @Description Create a new order_handler
// @Tags order_handler
// @Accept json
// @Produce json
// @Param order_handler body domain.Order true "Order"
// @Success 201 {object} domain.Order
// @Failure 400 {string} string 'StatusBadRequest'
// @Failure 500 {string} string 'StatusInternalServerError'
// @Router /order_handler [post]
func (h *OrderHandler) CreateOrder(c echo.Context) error {
	order := new(domain.Order)
	if err := c.Bind(order); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := h.service.CreateOrder(order); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, order)
}
