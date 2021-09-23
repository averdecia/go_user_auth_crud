package controllers

import (
	"encoding/json"

	"crud/database"
	"crud/entities"
)

var prefix = "sessions:"

// SaveSession function
func SaveSession(session entities.Session) (string, error) {

	key := prefix + session.Token
	data, err := json.Marshal(session)
	if err != nil {
		return key, err
	}

	_, err = database.Redis.Do("SET", key, data)
	if err != nil {
		return key, err
	}
	return key, nil
}
