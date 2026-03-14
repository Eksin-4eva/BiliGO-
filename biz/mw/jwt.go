package mw

import (
	"context"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/BiliGO/pkg/utils"
	"github.com/cloudwego/hertz/pkg/app"
)

const (
	CtxUserIDKey = "user_id"

	// 前端据此错误码发起 /user/refresh 请求
	CodeAccessExpired = 401001
	CodeTokenInvalid  = 401002
)

// JWTAuth 验证 Access Token，过期但 Refresh Token 有效时返回特定错误码
func JWTAuth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		secret := os.Getenv("JWT_SECRET")

		authHeader := string(c.GetHeader("Authorization"))
		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"code":    CodeTokenInvalid,
				"message": "missing or malformed token",
			})
			c.Abort()
			return
		}
		accessToken := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := utils.ParseToken(accessToken, secret)
		if err == utils.ErrTokenExpired {
			// Access Token 过期，检查 Refresh Token 是否有效
			refreshToken := string(c.GetHeader("X-Refresh-Token"))
			if refreshToken == "" {
				c.JSON(http.StatusUnauthorized, map[string]interface{}{
					"code":    CodeAccessExpired,
					"message": "access token expired, please refresh",
				})
				c.Abort()
				return
			}
			rClaims, rErr := utils.ParseToken(refreshToken, secret)
			if rErr != nil || rClaims.TokenType != utils.TokenTypeRefresh {
				c.JSON(http.StatusUnauthorized, map[string]interface{}{
					"code":    CodeAccessExpired,
					"message": "access token expired, refresh token invalid",
				})
				c.Abort()
				return
			}
			// Refresh Token 有效，告知前端刷新
			c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"code":    CodeAccessExpired,
				"message": "access token expired, please use refresh token to get new tokens",
				"user_id": strconv.FormatInt(rClaims.UserID, 10),
			})
			c.Abort()
			return
		}
		if err != nil {
			c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"code":    CodeTokenInvalid,
				"message": "invalid token",
			})
			c.Abort()
			return
		}
		if claims.TokenType != utils.TokenTypeAccess {
			c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"code":    CodeTokenInvalid,
				"message": "wrong token type",
			})
			c.Abort()
			return
		}

		c.Set(CtxUserIDKey, claims.UserID)
		c.Next(ctx)
	}
}
