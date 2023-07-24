package crud

import (
	"ege_platform/internal/logging"
	"ege_platform/internal/model"
	"fmt"

	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/models"
)

func CreateUser(dao *daos.Dao, claims *model.VKClaims, userVK api.UsersGetResponse) error {
	usersCollection, err := dao.FindCollectionByNameOrId("users")
	if err != nil {
		logging.Log.Errorf("Can't find collection 'users': %v", err)
		return err
	}

	newUser := models.NewRecord(usersCollection)
	newUser.Set("username", fmt.Sprintf("%d", claims.UserID))
	newUser.Set("name", userVK[0].FirstName+" "+userVK[0].LastName)
	newUser.Set("tokenKey", claims.AccessToken)

	return dao.SaveRecord(newUser)
}

func UpdateUserToken(dao *daos.Dao, claims *model.VKClaims, userPB *models.Record) error {
	userPB.Set("tokenKey", claims.AccessToken)
	return dao.SaveRecord(userPB)
}
