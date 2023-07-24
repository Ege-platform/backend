package controller

import (
	"ege_platform/internal/config"
	"ege_platform/internal/crud"
	"ege_platform/internal/logging"
	"ege_platform/internal/model"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/daos"
)

func GetVKClaims(c echo.Context, cfg *config.Config) (*model.VKClaims, error) {
	code := c.QueryParams().Get("code")

	response, err := http.Get(fmt.Sprintf(cfg.TokenURI, cfg.ClientID, cfg.ClientSecret, cfg.RedirectURI, code))
	if err != nil {
		logging.Log.Errorf("Can't get token: %v", err)
		return nil, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		logging.Log.Errorf("Error while reading response body: %v", err)
		return nil, err
	}

	claims := &model.VKClaims{}
	err = json.Unmarshal(body, &claims)
	if err != nil {
		logging.Log.Errorf("Can't parse json: %v", err)
		return nil, err
	}

	if claims.AccessToken == "" {
		logging.Log.Debug("Empty claims")
		return nil, errors.New("unauthorized")
	}

	return claims, nil
}

func GenerateLoginCredentials(c echo.Context, claims *model.VKClaims, cfg *config.Config) (*model.LoginCredentials, error) {
	vk := api.NewVK(claims.AccessToken)

	userVK, err := vk.UsersGet(api.Params{
		"user_ids": claims.UserID,
	})
	if err != nil {
		logging.Log.Errorf("Can't get user name from VK: %v", err)
		return nil, err
	}

	dao := c.Get("dao").(*daos.Dao)
	userPB, err := dao.FindAuthRecordByUsername("users", fmt.Sprintf("%d", claims.UserID))
	if err != nil {
		err = crud.CreateUser(dao, claims, userVK)
		if err != nil {
			logging.Log.Errorf("Can't create new user: %v", err)
			return nil, err
		}
	} else {
		err := crud.UpdateUserToken(dao, claims, userPB)
		if err != nil {
			logging.Log.Errorf("Can't update user's token: %v", err)
			return nil, err
		}
	}

	return &model.LoginCredentials{
		Username:    fmt.Sprintf("%d", claims.UserID),
		AccessToken: claims.AccessToken,
		Name:        userVK[0].FirstName + " " + userVK[0].LastName,
	}, nil

	// var passwd string

	// userPB, err := dao.FindAuthRecordByUsername("users", fmt.Sprintf("%d", claims.UserID))
	// if err != nil {
	// 	generatedPassword, err := password.Generate(16, 4, 4, true, false)
	// 	generatedPassword += fmt.Sprintf("%d", claims.UserID)
	// 	if err != nil {
	// 		logging.Log.Errorf("Can't generate password: %v", err)
	// 		return nil, err
	// 	}

	// 	block, err := aes.NewCipher([]byte(cfg.SecretKey))
	// 	if err != nil {
	// 		logging.Log.Errorf("Can't create cipher block: %v", err)
	// 		return nil, err
	// 	}

	// 	cipherText := make([]byte, aes.BlockSize+len(generatedPassword))
	// 	iv := cipherText[:aes.BlockSize]
	// 	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
	// 		logging.Log.Errorf("Can't read iv: %v", err)
	// 		return nil, err
	// 	}

	// 	stream := cipher.NewCFBEncrypter(block, iv)
	// 	stream.XORKeyStream(cipherText[aes.BlockSize:], []byte(generatedPassword))

	// 	passwd = base64.RawURLEncoding.EncodeToString(cipherText)
	// 	logging.Log.Debug(generatedPassword)

	// 	err = crud.CreateUser(dao, claims, userVK, passwd)
	// 	if err != nil {
	// 		logging.Log.Errorf("Can't create new user: %v", err)
	// 		return nil, err
	// 	}
	// } else {
	// 	passwd = userPB.PasswordHash()
	// }
}
