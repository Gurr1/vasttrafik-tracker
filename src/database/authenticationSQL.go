package database

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/gurr1/vasttrafik-tracker/entity"
	"time"
)

func CreateAuthentication(token string) {
	e := &entity.AuthenticationDetails{
		ID:             uuid.New().String(),
		Token:          token,
		UpdatedAt:      time.Now(),
		ExpirationTime: time.Now().Add(time.Hour * 2),
	}
	Database.Create(e)
}

func GetActiveAuthentication() (entity.AuthenticationDetails, error){
	var auth entity.AuthenticationDetails
	Database.First(&auth, "expiration_time > ?", time.Now())
	if &auth == nil {
		return entity.AuthenticationDetails{}, fmt.Errorf("no active authentication")
	}
	return auth, nil
	}

