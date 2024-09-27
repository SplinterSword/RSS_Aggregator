package database

import (
	"time"

	"github.com/google/uuid"
)

func (config *DBConfig) CreateFeed(user User, name string, url string) (any, error) {
	id := uuid.New().String()
	feed := Feed{
		ID:         id,
		Created_AT: time.Now(),
		Updated_AT: time.Now(),
		Name:       name,
		URL:        url,
		User_ID:    user.ID,
	}
	err := config.insertFeed(&feed)
	if err != nil {
		return Feed{}, err
	}

	id = uuid.New().String()
	feedfollow := FeedFollow{
		ID:         id,
		FeedID:     feed.ID,
		UserID:     user.ID,
		Created_At: time.Now(),
		Updated_AT: time.Now(),
	}

	err = config.insertFeedFollow(&feedfollow)
	if err != nil {
		return Feed{}, err
	}

	user.Feeds[url] = feed

	err = config.updateUser(user)
	if err != nil {
		return Feed{}, err
	}

	type response struct {
		Feed       Feed       `json:"feed"`
		FeedFollow FeedFollow `json:"feed_follow"`
	}

	return response{feed, feedfollow}, nil
}

func (config *DBConfig) GetAllFeeds() (any, error) {
	feeds, err := config.getFeeds()
	if err != nil {
		return nil, err
	}
	return feeds, nil
}
