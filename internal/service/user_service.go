package service

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"time"
	"tracking_service/internal/domain"
	"tracking_service/internal/repository"
	"tracking_service/pkg/authorization"
	"tracking_service/pkg/hash_password"
)

type UserService interface {
	Register(user domain.UserDTO) error
	Login(username, password string) (string, error)
	JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc
}

type userService struct {
	repo      repository.UserRepository
	jwtSecret string
}

func NewUserService(repo repository.UserRepository, jwtSecret string) UserService {
	return &userService{repo: repo, jwtSecret: jwtSecret}
}

func (s *userService) Register(user domain.UserDTO) error {
	hashedPassword, err := hash_password.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	return s.repo.Create(user)
}

func (s *userService) Login(username, password string) (string, error) {
	user, err := s.repo.GetByLogin(username)
	if err != nil {
		return "", errors.New("invalid login")
	}
	if !hash_password.IsValidPassword(password, user.Password) {
		return "", errors.New("invalid password")
	}
	token, err := authorization.CreateJWTToken(user.Login, s.jwtSecret)
	if err != nil {
		log.Fatal("чето с jwt")
		return "failed to create token", err
	}
	return token, nil
}

func (s *userService) generateJWT(user *domain.UserDTO) (string, error) {
	claims := jwt.MapClaims{
		"login":     user.Login,
		"isCourier": user.IsCourier,
		"exp":       time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.jwtSecret))
}

func (s *userService) JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				return c.JSON(http.StatusUnauthorized, echo.Map{"message": "missing or invalid token"})
			}
			return c.JSON(http.StatusBadRequest, echo.Map{"message": "bad request"})
		}

		tokenStr := cookie.Value
		claims, err := authorization.ValidateJWT(tokenStr, s.jwtSecret)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, echo.Map{"message": "unauthorized"})
		}

		c.Set("user", claims.Login)
		return next(c)
	}
}
