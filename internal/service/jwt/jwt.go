package service

import (
	"ege_platform/internal/crud"
	"ege_platform/internal/logging"
	"ege_platform/internal/model"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/daos"
)

func CreateAccessToken(username string, secret string) (string, error) {
	key := []byte(secret)

	jwtClaims := model.JwtClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			Subject:   username,
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)
	return accessToken.SignedString(key)
}

func VerifyAccessToken(accessTokenString string, secret string) (*jwt.Token, error) {
	accessToken, err := jwt.Parse(accessTokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

func AuthenticateUser(c echo.Context, claims *model.Claims, secret string) string {
	dao := c.Get("dao").(*daos.Dao)

	accessToken, err := CreateAccessToken(claims.Username, secret)
	if err != nil {
		logging.Log.Errorf("Can't create access token: %v", err)
		return ""
	}

	claims.AccessToken = accessToken

	userPB, exists := dao.FindAuthRecordByUsername("users", claims.Username)
	if exists != nil {
		err = crud.CreateUser(dao, claims)
		if err != nil {
			logging.Log.Errorf("Can't create user record: %v", err)
			return ""
		}
	}

	err = crud.UpdateUserToken(dao, claims, userPB)
	if err != nil {
		logging.Log.Errorf("Can't update user token: %v", err)
		return ""
	}

	return accessToken
}
