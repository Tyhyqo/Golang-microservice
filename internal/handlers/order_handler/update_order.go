package order_handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"tracking_service/internal/domain"
)

// UpdateOrder godoc
// @Summary Update an order_handler by ID
// @Description Update an order_handler by ID
// @Tags order_handler
// @Accept json
// @Produce json
// @Param id path int true "Order ID"
// @Param order_handler body domain.Order true "Order"
// @Success 201 {object} domain.Order
// @Failure 400 {string} string 'StatusBadRequest'
// @Failure 500 {string} string 'StatusInternalServerError'
// @Router /order_handler/{id} [put]
func (h *OrderHandler) UpdateOrder(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	order := new(domain.Order)
	if err := c.Bind(order); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	order.ID = uint(id)

	if err := h.service.UpdateOrder(order); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, order)
}
