package auth

import (
	"fmt"
	"github.com/abishz17/go-backend-template/infrastructure"
	"github.com/abishz17/go-backend-template/internal/domain"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"time"
)

const (
	ACCESS_TOKEN_VALIDITY_PERIOD  time.Duration = 60 * time.Minute
	REFRESH_TOKEN_VALIDITY_PERIOD time.Duration = 7 * 24 * 60 * time.Minute
)

type Token struct {
	AccessToken  string
	RefreshToken string
}

type TokenClaims struct {
	UserId       uuid.UUID              `json:"user_id"`
	CustomClaims map[string]interface{} `json:"custom_claims"`
	jwt.StandardClaims
}

func GenerateJWTTokens(user domain.User, customClaims map[string]interface{}, env infrastructure.Env) (*Token, error) {
	accessTokenKey := env.AccessTokenSecret
	refreshTokenKey := env.RefreshTokenSecret

	accessTokenClaims := TokenClaims{
		UserId:       user.ID,
		CustomClaims: customClaims,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(ACCESS_TOKEN_VALIDITY_PERIOD).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	refreshTokenClaims := TokenClaims{
		UserId:       user.ID,
		CustomClaims: customClaims,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(REFRESH_TOKEN_VALIDITY_PERIOD).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)

	accessTokenString, accessErr := accessToken.SignedString([]byte(accessTokenKey))
	refreshTokenString, refreshErr := refreshToken.SignedString([]byte(refreshTokenKey))
	if accessErr != nil || refreshErr != nil {
		return nil, fmt.Errorf("error generating tokens")
	}
	return &Token{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}, nil

}

//func VerifyToken(token string)
