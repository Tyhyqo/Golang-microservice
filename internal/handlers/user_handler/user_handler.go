package user_handler

import "tracking_service/internal/service"

// UserHandler handles user_handler related requests
type UserHandler struct {
	service service.UserService
}

// NewUserHandler creates a new UserHandler
func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{service: service}
}
