package entity

import "github.com/golang/protobuf/ptypes/timestamp"

type AuthenticationDetails struct {
	Id				string					`json:"id"`
	Token 			string					`json:"token"`
	updatedAt		timestamp.Timestamp
	expirationTime	int
}
