package domain

import "github.com/dgrijalva/jwt-go"

type LoginDTO struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Active    bool   `json:"active"`
	CreatedAt string `json:"createdAt"`
	jwt.StandardClaims
}
