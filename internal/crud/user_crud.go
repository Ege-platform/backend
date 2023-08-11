package crud

import (
	"ege_platform/internal/logging"
	"ege_platform/internal/model"

	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/models"
)

func CreateUser(dao *daos.Dao, claims *model.Claims) (*model.User, error) {
	usersCollection, err := dao.FindCollectionByNameOrId("users")
	if err != nil {
		logging.Log.Errorf("Can't find collection 'users': %v", err)
		return nil, err
	}

	newUser := models.NewRecord(usersCollection)
	newUser.Set("username", claims.Username)
	newUser.Set("name", claims.Name)
	newUser.Set("tokenKey", claims.AccessToken)
	newUser.Set("coins", 0)
	newUser.Set("experience", 0)

	err = dao.SaveRecord(newUser)
	if err != nil {
		logging.Log.Errorf("Can't save user to PB: %v", err)
		return nil, err
	}

	return model.RecordToModel(newUser), nil
}

func UpdateUserToken(dao *daos.Dao, claims *model.Claims, userPB *models.Record) (*model.User, error) {
	userPB.Set("tokenKey", claims.AccessToken)
	err := dao.SaveRecord(userPB)
	if err != nil {
		logging.Log.Errorf("Can't update token for user record: %v", err)
		return nil, err
	}

	return model.RecordToModel(userPB), nil
}
