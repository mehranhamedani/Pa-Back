package repository

import (
	"context"
	"fmt"
	"log"
	"pa-back/model"

	"go.mongodb.org/mongo-driver/bson"
)

func getOrCreateUserByMobileNumber(mobileNumber string) (*model.User, error) {
	var error error
	userModel := model.User{}
	usersCollection := db.Collection("users")
	ctx := context.TODO()
	result := usersCollection.FindOne(ctx, bson.M{"firstName": "Mehran"})
	error = result.Decode(&userModel)
	//cur.Decode(userModel)
	// for cur.Next(ctx) {
	// 	var result bson.M
	// 	err := cur.Decode(&result)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	// do something with result....
	// }
	fmt.Printf(result.Err().Error())
	singleResult := usersCollection.FindOne(ctx, bson.M{"mobileNumber": mobileNumber})
	if singleResult.Err() != nil {
		log.Println(singleResult.Err().Error())
	}
	error = singleResult.Decode(&userModel)
	if error != nil {
		log.Println(error.Error())
	}
	return &userModel, error
}
