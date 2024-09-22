package database

import (
	"crypto/rand"
	"encoding/base64"
	"time"

	"github.com/google/uuid"
)

func generateApiKey() (string, error) {
	apiKeyBytes := make([]byte, 64)
	_, err := rand.Read(apiKeyBytes)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(apiKeyBytes), nil
}

func (config *DBConfig) MakeUser(name string) (User, error) {

	id := uuid.New().String()
	api_key, err := generateApiKey()
	if err != nil {
		return User{}, err
	}

	user := User{
		ID:         id,
		Created_AT: time.Now(),
		Updated_AT: time.Now(),
		Name:       name,
		Api_Key:    api_key,
		Feeds:      make(map[string]Feed),
	}

	err = config.insertUser(&user)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (config *DBConfig) DeleteUser(id string) error {
	err := config.deleteUser(id)
	if err != nil {
		return err
	}
	return nil
}

func (config *DBConfig) UpdateUser(user User) error {
	err := config.updateUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (config *DBConfig) GetUserByApiKey(api_key string) (User, error) {
	user, err := config.getUserbyApiKey(api_key)
	if err != nil {
		return User{}, err
	}
	return user, nil
}
