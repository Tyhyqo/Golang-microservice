package http

import (
	"net/http"
	"tracking_service/internal/domain"
	"tracking_service/internal/service"

	"github.com/labstack/echo/v4"
)

// UserHandler handles user related requests
type UserHandler struct {
	service service.UserService
}

// NewUserHandler creates a new UserHandler
func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

// RegisterUser godoc
// @Summary Register a new user
// @Description Register a new user
// @Tags users
// @Accept json
// @Produce json
// @Param user body domain.User true "User"
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

// LoginUser godoc
// @Summary Login a user
// @Description Login a user
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
