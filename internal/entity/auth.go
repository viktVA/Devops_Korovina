package entity

import (
	"github.com/golang-jwt/jwt"
	"time"
)

type Claims struct {
	UserID    UserID    `json:"user_id"`
	Login     string    `json:"login_name"`
	Nickname  string    `json:"nickname"`
	ExpiresAt time.Time `json:"expiresAt"`
	jwt.StandardClaims
}
