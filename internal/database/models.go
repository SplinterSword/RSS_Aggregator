package database

import "time"

type User struct {
	ID         string          `bson:"_id" validate:"required"`
	Created_AT time.Time       `bson:"created_at" validate:"required"`
	Updated_AT time.Time       `bson:"updated_at" validate:"required"`
	Name       string          `bson:"name" validate:"required"`
	Api_Key    string          `bson:"api_key" validate:"required"`
	Feeds      map[string]Feed `bson:"feeds"`
}

type Feed struct {
	ID          string    `bson:"_id" validate:"required"`
	Created_AT  time.Time `bson:"created_at" validate:"required"`
	Updated_AT  time.Time `bson:"updated_at" validate:"required"`
	Name        string    `bson:"name" validate:"required"`
	URL         string    `bson:"url" validate:"required"`
	User_ID     string    `bson:"user_id" validate:"required"`
	LastFetched time.Time `bson:"last_fetched_at"`
}

type FeedFollow struct {
	ID         string    `bson:"_id" validate:"required"`
	FeedID     string    `bson:"feed_id" validate:"required"`
	UserID     string    `bson:"user_id" validate:"required"`
	Created_At time.Time `bson:"created_at" validate:"required"`
	Updated_AT time.Time `bson:"updated_at" validate:"required"`
}
