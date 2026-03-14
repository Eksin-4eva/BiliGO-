package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	AccessTokenDuration  = 2 * time.Hour
	RefreshTokenDuration = 7 * 24 * time.Hour

	TokenTypeAccess  = "access"
	TokenTypeRefresh = "refresh"
)

var (
	ErrTokenExpired   = errors.New("token expired")
	ErrTokenInvalid   = errors.New("token invalid")
	ErrTokenTypeMismatch = errors.New("token type mismatch")
)

type Claims struct {
	UserID    int64  `json:"user_id"`
	TokenType string `json:"token_type"`
	jwt.RegisteredClaims
}

// GenerateAccessToken 生成 Access Token，有效期 2 小时
func GenerateAccessToken(userID int64, secret string) (string, error) {
	return generateToken(userID, TokenTypeAccess, AccessTokenDuration, secret)
}

// GenerateRefreshToken 生成 Refresh Token，有效期 7 天
func GenerateRefreshToken(userID int64, secret string) (string, error) {
	return generateToken(userID, TokenTypeRefresh, RefreshTokenDuration, secret)
}

func generateToken(userID int64, tokenType string, duration time.Duration, secret string) (string, error) {
	claims := Claims{
		UserID:    userID,
		TokenType: tokenType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secret))
}

// ParseToken 解析 Token，返回 Claims
// 过期时也会返回 claims，调用方可据此判断是否需要刷新
func ParseToken(tokenStr, secret string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrTokenInvalid
		}
		return []byte(secret), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return claims, ErrTokenExpired
		}
		return nil, ErrTokenInvalid
	}
	if !token.Valid {
		return nil, ErrTokenInvalid
	}
	return claims, nil
}
