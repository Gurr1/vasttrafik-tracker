package entity

import (
	"time"
)

type AuthenticationDetails struct {
	ID             string    `json:"id"`
	Token          string    `json:"token"`
	UpdatedAt      time.Time `json:"updated_at"`
	ExpirationTime time.Time `json:"expiration_time"`
}



func IsExpired(details AuthenticationDetails) bool{

	return false
}