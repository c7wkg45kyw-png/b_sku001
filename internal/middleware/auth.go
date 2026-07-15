package middleware

import (
	"net/http"
	"strings"

	"bsku001/backend/internal/config"
	"bsku001/backend/internal/model"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

const AuthContextKey = "auth_context"

type Claims struct {
	MerchantID string   `json:"merchant_id"`
	ClientID   string   `json:"client_id"`
	Type       string   `json:"type"`
	Scopes     []string `json:"scopes"`
	jwt.RegisteredClaims
}

func Auth(cfg config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if !strings.HasPrefix(header, "Bearer ") {
			abort(c, http.StatusUnauthorized, "UNAUTHORIZED", "missing bearer token")
			return
		}
		tokenText := strings.TrimSpace(strings.TrimPrefix(header, "Bearer "))
		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenText, claims, func(token *jwt.Token) (any, error) { return []byte(cfg.JWTSecret), nil }, jwt.WithIssuer(cfg.JWTIssuer), jwt.WithAudience(cfg.JWTAudience))
		if err != nil || !token.Valid {
			abort(c, http.StatusUnauthorized, "UNAUTHORIZED", "invalid token")
			return
		}
		merchantID := claims.MerchantID
		if merchantID == "" {
			merchantID = c.GetHeader("X-Merchant-ID")
		}
		if merchantID == "" {
			abort(c, http.StatusUnauthorized, "UNAUTHORIZED", "merchant_id claim is required")
			return
		}
		c.Set(AuthContextKey, model.AuthContext{UserID: claims.Subject, MerchantID: merchantID, Scopes: claims.Scopes})
		c.Next()
	}
}

func RequireScope(scope string) gin.HandlerFunc {
	return func(c *gin.Context) {
		auth, _ := c.Get(AuthContextKey)
		ctx, ok := auth.(model.AuthContext)
		if !ok {
			abort(c, http.StatusUnauthorized, "UNAUTHORIZED", "missing auth context")
			return
		}
		for _, current := range ctx.Scopes {
			if current == scope || current == "sku:*" {
				c.Next()
				return
			}
		}
		abort(c, http.StatusForbidden, "FORBIDDEN", "missing scope: "+scope)
	}
}

func CurrentAuth(c *gin.Context) model.AuthContext {
	value, _ := c.Get(AuthContextKey)
	ctx, _ := value.(model.AuthContext)
	return ctx
}

func abort(c *gin.Context, status int, code, message string) {
	c.AbortWithStatusJSON(status, model.ErrorResponse{Success: false, Code: code, Message: message})
}
