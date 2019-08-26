package repository

import (
	"context"
	"fmt"
	"log"
	"pa-back/model"
	"pa-back/resources/enums"
	"pa-back/resources/texts"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/bson"
)

// GetOrCreateUserByMobileNumber func
func GetOrCreateUserByMobileNumber(mobileNumber string) (*model.User, error) {
	userModel, error := GetUserByMobileNumber(mobileNumber)
	if userModel.UserID.IsZero() {
		userModel, error = CreateUser(mobileNumber)
	}
	return userModel, error
}

// GetUserByUserID func
func GetUserByUserID(userID primitive.ObjectID) (*model.User, error) {
	var error error
	userModel := model.User{}
	usersCollection := db.Collection(enums.DBCollectionNameUsers.ToString())
	result := usersCollection.FindOne(context.Background(), bson.M{"_id": userID})
	if result.Err() != nil && result.Err().Error() != texts.En_MNDIR {
		log.Println(result.Err())
		error = result.Err()
	}
	error = result.Decode(&userModel)
	return &userModel, error
}

// GetUserByMobileNumber func
func GetUserByMobileNumber(mobileNumber string) (*model.User, error) {
	var error error
	userModel := model.User{}
	usersCollection := db.Collection(enums.DBCollectionNameUsers.ToString())
	result := usersCollection.FindOne(context.Background(), bson.M{"mobileNumber": mobileNumber})
	if result.Err() != nil && result.Err().Error() != texts.En_MNDIR {
		log.Println(result.Err())
		error = result.Err()
	}
	decodeError := result.Decode(&userModel)
	if decodeError != nil && decodeError.Error() != texts.En_MNDIR {
		log.Println(decodeError)
		error = decodeError
	}
	return &userModel, error
}

// CreateUser func
func CreateUser(mobileNumber string) (*model.User, error) {
	var error error
	userModel := &model.User{}
	usersCollection := db.Collection(enums.DBCollectionNameUsers.ToString())
	res, err := usersCollection.InsertOne(context.Background(), bson.M{"mobileNumber": mobileNumber})
	if err != nil {
		log.Println(err)
		error = err
	} else {
		userModel, error = GetUserByUserID(res.InsertedID.(primitive.ObjectID))
	}
	return userModel, error
}

// GetAllUsers func
func GetAllUsers() {
	fmt.Println("")

	//cur.Decode(userModel)
	// for cur.Next(ctx) {
	// 	var result bson.M
	// 	err := cur.Decode(&result)
	// 	if err != nil {
	// 		log.Println(err)
	// 	}
	// 	// do something with result....
	// }
}
