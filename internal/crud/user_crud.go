package crud

import (
	"ege_platform/internal/logging"
	"ege_platform/internal/model"
	"fmt"

	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/models"
)

func CreateUser(dao *daos.Dao, claims *model.Claims) error {
	usersCollection, err := dao.FindCollectionByNameOrId("users")
	if err != nil {
		logging.Log.Errorf("Can't find collection 'users': %v", err)
		return err
	}

	newUser := models.NewRecord(usersCollection)
	newUser.Set("username", fmt.Sprintf("%d", claims.ID))
	newUser.Set("name", claims.Name)
	newUser.Set("tokenKey", claims.AccessToken)

	return dao.SaveRecord(newUser)
}

func UpdateUserToken(dao *daos.Dao, claims *model.Claims, userPB *models.Record) error {
	userPB.Set("tokenKey", claims.AccessToken)
	return dao.SaveRecord(userPB)
}
