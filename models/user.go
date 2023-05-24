package models

import (
	"github.com/vinitparekh17/go-mongo/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	EMAIL    string             `bson:"email"`
	PASSWORD string             `bson:"password"`
}

func (u *User) GetByID(id string) (User, error) {
	var user User
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return user, err
	}
	err = db.Collection.FindOne(db.Ctx, bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (u *User) AddUser(user User) (interface{}, error) {
	result, err := db.Collection.InsertOne(db.Ctx, user)
	if err != nil {
		return nil, err
	}
	return result.InsertedID, nil
}
