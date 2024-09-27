package database

import (
	"time"

	"github.com/google/uuid"
)

func (config *DBConfig) CreateFeedFollows(user User, feedID string) (FeedFollow, error) {

	feed, err := config.getFeedbyFeedID(feedID)
	if err != nil {
		return FeedFollow{}, err
	}

	feedFollow := FeedFollow{
		ID:         uuid.New().String(),
		FeedID:     feed.ID,
		UserID:     user.ID,
		Created_At: time.Now(),
		Updated_AT: time.Now(),
	}
	err = config.insertFeedFollow(&feedFollow)
	if err != nil {
		return FeedFollow{}, err
	}
	return feedFollow, nil
}

func (config *DBConfig) GetAllFeedFollows() ([]FeedFollow, error) {
	feedFollows, err := config.getFeedFollows()
	if err != nil {
		return nil, err
	}
	return feedFollows, nil
}

func (config *DBConfig) DeleteFeedFollows(user User, feedID string) error {
	feed, err := config.getFeedbyFeedID(feedID)
	if err != nil {
		return err
	}
	err = config.deleteFeedFollows(user, feed)
	if err != nil {
		return err
	}
	return nil
}
