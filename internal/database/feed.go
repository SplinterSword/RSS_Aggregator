package database

import (
	"time"

	"github.com/google/uuid"
)

func (config *DBConfig) CreateFeed(user User, name string, url string) (Feed, error) {
	id := uuid.New().String()
	feed := Feed{
		ID:         id,
		Created_AT: time.Now(),
		Updated_AT: time.Now(),
		Name:       name,
		URL:        url,
		User_ID:    user.ID,
	}

	user.Feeds[url] = feed

	err := config.updateUser(user)
	if err != nil {
		return Feed{}, err
	}

	return feed, nil
}
