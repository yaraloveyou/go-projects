package models

import (
	"context"
	"errors"
	"go-mongodb/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type User struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Name    string             `bson:"name"`
	Surname string             `bson:"surname"`
	Age     int                `bson:"age"`
	City    string             `bson:"city"`
}

func (u *User) CreateUser() (*User, error) {
	_, err := db.Collection().InsertOne(context.TODO(), u)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (u *User) UpdateUser() (*User, error) {
	filter := bson.M{"_id": u.ID}

	update := bson.M{"$set": bson.M{
		"name":    u.Name,
		"surname": u.Surname,
		"age":     u.Age,
		"city":    u.City,
	}}

	result, err := db.Collection().UpdateOne(context.TODO(), filter, update)

	if err != nil {
		return nil, err
	}

	if result.ModifiedCount == 0 {
		return nil, errors.New("user not found")
	}
	return u, nil
}

func DeleteUsers(id string) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objectId}

	_, err = db.Collection().DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}

	return nil
}

func GetUsers() ([]*User, error) {
	cur, err := db.Collection().Find(context.TODO(), bson.M{}, options.Find())
	if err != nil {
		log.Fatal(err)
	}
	var users []*User
	for cur.Next(context.TODO()) {
		var elem User
		err := cur.Decode(&elem)
		if err != nil {
			return nil, err
		}

		users = append(users, &elem)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	cur.Close(context.TODO())
	return users, nil
}

func GetUser(id string) (*User, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": objectID}
	var elem User
	err = db.Collection().FindOne(context.TODO(), filter).Decode(&elem)
	if err != nil {
		return nil, err
	}

	return &elem, nil
}
