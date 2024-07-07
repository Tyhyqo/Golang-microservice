package order_handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

// GetOrder godoc
// @Summary Get an order_handler by ID
// @Description Get an order_handler by ID
// @Tags order_handler
// @Accept json
// @Produce json
// @Param id path int true "Order ID"
// @Success 201 {object} domain.Order
// @Failure 400 {string} string 'StatusBadRequest'
// @Failure 500 {string} string 'StatusInternalServerError'
// @Router /order_handler/{id} [get]
func (h *OrderHandler) GetOrder(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	order, err := h.service.GetOrderById(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, order)
}
