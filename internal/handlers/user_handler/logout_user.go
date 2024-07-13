package user_handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"tracking_service/pkg/authorization"
)

// LogoutUser godoc
// @Summary Logout the user
// @Description Logout the user by clearing the JWT cookie
// @Tags users
// @Success 200 {string} string "logged out successfully"
// @Router /logout [post]
func (h *UserHandler) LogoutUser(c echo.Context) error {
	authorization.ClearJWTCookie(c)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "logged out successfully",
	})
}
