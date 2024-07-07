package http

import (
	"net/http"
	"strconv"
	"tracking_service/internal/domain"
	"tracking_service/internal/service"

	"github.com/labstack/echo/v4"
)

type OrderHandler struct {
	service service.OrderService
}

func NewOrderHandler(service service.OrderService) *OrderHandler {
	return &OrderHandler{service: service}
}

// CreateOrder godoc
// @Summary Create a new order
// @Description Create a new order
// @Tags orders
// @Accept json
// @Produce json
// @Param order body domain.Order true "Order"
// @Success 201 {object} domain.Order
// @Failure 400 {string} string 'StatusBadRequest'
// @Failure 500 {string} string 'StatusInternalServerError'
// @Router /orders [post]
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

// GetOrder godoc
// @Summary Get an order by ID
// @Description Get an order by ID
// @Tags orders
// @Accept json
// @Produce json
// @Param id path int true "Order ID"
// @Success 201 {object} domain.Order
// @Failure 400 {string} string 'StatusBadRequest'
// @Failure 500 {string} string 'StatusInternalServerError'
// @Router /orders/{id} [get]
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

// UpdateOrder godoc
// @Summary Update an order by ID
// @Description Update an order by ID
// @Tags orders
// @Accept json
// @Produce json
// @Param id path int true "Order ID"
// @Param order body domain.Order true "Order"
// @Success 201 {object} domain.Order
// @Failure 400 {string} string 'StatusBadRequest'
// @Failure 500 {string} string 'StatusInternalServerError'
// @Router /orders/{id} [put]
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
