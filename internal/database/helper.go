package database

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (config *DBConfig) insertUser(user *User) error {
	client := config.client
	database := client.Database("Blogator")
	collection := database.Collection("users")

	_, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}

	return nil
}

func (config *DBConfig) insertFeed(feed *Feed) error {
	client := config.client
	database := client.Database("Blogator")
	collection := database.Collection("feeds")

	_, err := collection.InsertOne(context.TODO(), feed)
	if err != nil {
		return err
	}

	return nil
}

func (config *DBConfig) insertFeedFollow(feedFollow *FeedFollow) error {
	client := config.client
	database := client.Database("Blogator")
	collection := database.Collection("feed_follows")

	_, err := collection.InsertOne(context.TODO(), feedFollow)
	if err != nil {
		return fmt.Errorf("error inserting feed_follow: %v", err)
	}

	return nil
}

func (config *DBConfig) getUserbyApiKey(api_key string) (User, error) {
	client := config.client
	database := client.Database("Blogator")
	collection := database.Collection("users")

	var filter struct {
		Api_Key string `bson:"api_key"`
	}

	filter.Api_Key = api_key

	user := User{}
	err := collection.FindOne(context.Background(), filter).Decode(&user)

	if err != nil {
		return User{}, errors.New("User not found")
	}

	return user, nil
}
func (config *DBConfig) getFeedbyFeedID(feed_id string) (Feed, error) {
	client := config.client
	database := client.Database("Blogator")
	collection := database.Collection("feeds")

	var filter struct {
		feed_id string `bson:"api_key"`
	}

	filter.feed_id = feed_id

	feed := Feed{}
	err := collection.FindOne(context.Background(), filter).Decode(&feed)

	if err != nil {
		return Feed{}, errors.New("Feed not found")
	}

	return feed, nil
}

func (config *DBConfig) getFeedFollows() ([]FeedFollow, error) {
	client := config.client
	database := client.Database("Blogator")
	collection := database.Collection("feed_follows")

	feedfollows := []FeedFollow{}

	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, fmt.Errorf("error finding feed_follows: %v", err)
	}
	defer cursor.Close(context.Background())

	err = cursor.All(context.Background(), &feedfollows)

	if err != nil {
		return nil, fmt.Errorf("error decoding feed_follows: %v", err)
	}
	return feedfollows, nil

}

func (config *DBConfig) getFeeds() (any, error) {
	client := config.client
	database := client.Database("Blogator")
	collection := database.Collection("users")

	// Correct projection: Exclude everything except "feeds"
	projection := bson.D{
		{"_id", 0},   // Exclude the "_id" field
		{"feeds", 1}, // Include the "feeds" field
	}

	opts := options.Find().SetProjection(projection)

	cursor, err := collection.Find(context.Background(), bson.M{}, opts)
	if err != nil {
		return nil, fmt.Errorf("error finding feeds: %v", err)
	}
	defer cursor.Close(context.Background())

	var feeds []any
	for cursor.Next(context.Background()) {
		var feed bson.M
		err := cursor.Decode(&feed)
		if err != nil {
			return nil, fmt.Errorf("error decoding feed: %v", err)
		}

		// Ensure "feeds" field exists
		if f, ok := feed["feeds"]; ok {
			feeds = append(feeds, f)
		} else {
			return nil, errors.New("feeds field not found in document")
		}
	}

	// Check for any cursor errors
	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("error iterating through cursor: %v", err)
	}

	return feeds, nil
}

func (config *DBConfig) updateUser(user User) error {
	client := config.client
	database := client.Database("Blogator")
	collection := database.Collection("users")

	var filter struct {
		ID string `bson:"_id"`
	}

	filter.ID = user.ID

	update := bson.M{
		"$set": bson.M{
			"name":       user.Name,
			"updated_at": time.Now(),
			"feeds":      user.Feeds,
		},
	}

	log.Println(user)

	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (config *DBConfig) deleteUser(id string) error {
	client := config.client
	database := client.Database("Blogator")
	collection := database.Collection("users")

	_, err := collection.DeleteOne(context.TODO(), User{ID: id})
	if err != nil {
		return err
	}
	return nil
}

func (config *DBConfig) deleteFeedFollows(user User, feed Feed) error {
	client := config.client
	database := client.Database("Blogator")
	collection := database.Collection("feed_follows")

	var filter struct {
		UserID string `bson:"user_id"`
		FeedID string `bson:"feed_id"`
	}

	filter.FeedID = feed.ID
	filter.UserID = user.ID

	_, err := collection.DeleteMany(context.TODO(), filter)
	if err != nil {
		return fmt.Errorf("error deleting feed_follows: %v", err)
	}
	return nil
}
