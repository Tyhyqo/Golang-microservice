package domain

import "github.com/golang-jwt/jwt"

type UserDTO struct {
	Login     string `gorm:"unique;not null;foreignKey"`
	Password  string `gorm:"not null"`
	IsCourier bool   `gorm:"not null"`
}

type UserWeb struct {
	Login     string `json:"login"`
	Password  string `json:"hash_password"`
	IsCourier bool   `json:"isCourier"`
}

type ClaimsUser struct {
	Login string `json:"login"`
	jwt.StandardClaims
}
