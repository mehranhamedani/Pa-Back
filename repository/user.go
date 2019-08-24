package repository

import (
	"context"
	"fmt"
	"log"
	"pa-back/model"
	"pa-back/resources/enums"
	"pa-back/resources/texts"

	"go.mongodb.org/mongo-driver/bson"
)

func getOrCreateUserByMobileNumber(mobileNumber string) (*model.User, error) {
	var error error
	userModel := model.User{}
	usersCollection := db.Collection(enums.DBCollectionName_Users.ToString())
	result := usersCollection.FindOne(context.Background(), bson.M{"mobileNumber": mobileNumber})
	if result.Err() != nil && result.Err().Error() != texts.En_MNDIR {
		log.Println(result.Err())
		return &userModel, result.Err()
	}
	error = result.Decode(&userModel)
	if error != nil && error.Error() != texts.En_MNDIR {
		log.Println(error)
		return &userModel, error
	}
	if userModel.UserId == "" {
		res, err := usersCollection.InsertOne(context.Background(), bson.M{"mobileNumber": mobileNumber})
		if err != nil {
			log.Println(err)
			return &userModel, err
		}

	}
	return &userModel, error
}

func getUserByUserId(userId string) (*model.User, error) {
	var error error
	userModel := model.User{}
	usersCollection := db.Collection(enums.DBCollectionName_Users.ToString())
	result := usersCollection.FindOne(context.Background(), bson.M{"_id": userId})
	if result.Err() != nil && result.Err().Error() != texts.En_MNDIR {
		log.Println(result.Err())
		error = result.Err()
	}
	error = result.Decode(&userModel)
	return &userModel, error
}

func createUser(mobileNumber string) (*model.User, error) {
	var error error
	userModel := model.User{}
	usersCollection := db.Collection(enums.DBCollectionName_Users.ToString())
	result := usersCollection.FindOne(context.Background(), bson.M{"mobileNumber": mobileNumber})
	if result.Err() != nil && result.Err().Error() != texts.En_MNDIR {
		log.Println(result.Err())
		return &userModel, result.Err()
	}
	error = result.Decode(&userModel)
	if error != nil && error.Error() != texts.En_MNDIR {
		log.Println(error)
		return &userModel, error
	}
	if userModel.UserId == "" {

	}
	return &userModel, error
}

func getAllUsers() {
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
