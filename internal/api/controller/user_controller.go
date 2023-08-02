package controller

import (
	"ege_platform/internal/crud"
	"ege_platform/internal/model"
	"fmt"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/daos"
)

func CreateUser(c echo.Context, claims *model.Claims) error {
	dao := c.Get("dao").(*daos.Dao)
	userPB, err := dao.FindAuthRecordByUsername("users", fmt.Sprintf("%d", claims.ID))
	if err != nil {
		return crud.CreateUser(dao, claims)
	}
	return crud.UpdateUserToken(dao, claims, userPB)
}
