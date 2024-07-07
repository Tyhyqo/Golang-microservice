package user_handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"tracking_service/internal/domain"
)

// RegisterUser godoc
// @Summary Register a new user_handler
// @Description Register a new user_handler
// @Tags users
// @Accept json
// @Produce json
// @Param user_handler body domain.User true "User"
// @Success 201 {object} domain.User
// @Failure 400 {string} string 'StatusBadRequest'
// @Failure 500 {string} string 'StatusInternalServerError'
// @Router /users/register [post]
func (h *UserHandler) RegisterUser(c echo.Context) error {
	user := new(domain.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := h.service.Register(user); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, user)
}
