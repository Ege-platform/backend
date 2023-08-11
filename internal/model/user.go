package model

import (
	"time"

	"github.com/pocketbase/pocketbase/models"
)

type User struct {
	ID              string    `json:"id"`
	CollectionID    string    `json:"collectionId"`
	Username        string    `json:"username"`
	Verified        bool      `json:"verified"`
	EmailVisibility bool      `json:"emailVisibility"`
	Email           string    `json:"email"`
	Created         time.Time `json:"created"`
	Updated         time.Time `json:"updated"`
	Name            string    `json:"name"`
	Coins           float64   `json:"coins"`
	Experience      float64   `json:"experience"`
}

func RecordToModel(userRecord *models.Record) *User {
	return &User{
		ID:              userRecord.Id,
		CollectionID:    userRecord.Collection().Id,
		Username:        userRecord.Username(),
		Verified:        userRecord.Verified(),
		EmailVisibility: userRecord.EmailVisibility(),
		Email:           userRecord.Email(),
		Created:         userRecord.Created.Time(),
		Updated:         userRecord.Updated.Time(),
		Name:            userRecord.Get("name").(string),
		Coins:           userRecord.Get("coins").(float64),
		Experience:      userRecord.Get("experience").(float64),
	}
}

/*

{
  "token": "str",
  "model":{
    "id": "RECORD_ID",
      "collectionId": "_pb_users_auth_",
      "collectionName": "users",
      "username": "username123",
      "verified": false,
      "emailVisibility": true,
      "email": "test@example.com",
      "created": "2022-01-01 01:00:00.123Z",
      "updated": "2022-01-01 23:59:59.456Z",
      "name": "test",
      "coins": 123,
      "experience": 123
    }
  }
*/
