package database

import (
	"context"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
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
