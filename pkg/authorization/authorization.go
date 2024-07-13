package authorization

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
	"tracking_service/internal/domain"
)

func CreateJWTToken(login, jwtSecret string) (string, error) {
	claims := &domain.ClaimsUser{
		Login: login,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecret))
}

func SetJWTCookie(c echo.Context, token string) {
	cookie := &http.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
		Secure:   true,
	}
	c.SetCookie(cookie)
}

func ClearJWTCookie(c echo.Context) {
	cookie := &http.Cookie{
		Name:     "token",
		Value:    "",
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
		Secure:   true,
	}
	c.SetCookie(cookie)
}

func ValidateJWT(tokenString, jwtSecret string) (*domain.ClaimsUser, error) {
	claims := &domain.ClaimsUser{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, err
	}

	return claims, nil
}
