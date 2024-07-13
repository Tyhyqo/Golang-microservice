package user_handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"tracking_service/internal/domain"
	"tracking_service/pkg/authorization"
)

// LoginUser godoc
// @Summary Login a user_handler
// @Description Login a user_handler
// @Tags users
// @Accept json
// @Produce json
// @Param user_handler body domain.UserWeb true "UserWeb"
// @Success 200 {object} map[string]string
// @Failure 400 {string} string 'StatusBadRequest'
// @Failure 500 {string} string 'StatusInternalServerError'
// @Router /users/login [post]
func (h *UserHandler) LoginUser(c echo.Context) error {
	credentials := new(domain.UserWeb)
	if err := c.Bind(&credentials); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	token, err := h.service.Login(credentials.Login, credentials.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, err)
	}
	authorization.SetJWTCookie(c, token)

	return c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}
