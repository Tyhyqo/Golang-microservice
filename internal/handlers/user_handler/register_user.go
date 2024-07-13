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
// @Param user_handler body domain.UserWeb true "UserWeb"
// @Success 201 {object} domain.UserWeb
// @Failure 400 {string} string 'StatusBadRequest'
// @Failure 500 {string} string 'StatusInternalServerError'
// @Router /users/register [post]
func (h *UserHandler) RegisterUser(c echo.Context) error {
	user := new(domain.UserWeb)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	newUser := domain.UserDTO{
		Login:     user.Login,
		Password:  user.Password,
		IsCourier: user.IsCourier,
	}

	if err := h.service.Register(newUser); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, user)
}
