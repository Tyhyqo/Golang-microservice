package user_handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// LoginUser godoc
// @Summary Login a user_handler
// @Description Login a user_handler
// @Tags users
// @Accept json
// @Produce json
// @Param credentials body map[string]string true "Credentials"
// @Success 200 {object} map[string]string
// @Failure 400 {string} string 'StatusBadRequest'
// @Failure 500 {string} string 'StatusInternalServerError'
// @Router /users/login [post]
func (h *UserHandler) LoginUser(c echo.Context) error {
	credentials := make(map[string]string)
	if err := c.Bind(&credentials); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	token, err := h.service.Login(credentials["username"], credentials["password"])
	if err != nil {
		return c.JSON(http.StatusUnauthorized, err)
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}
