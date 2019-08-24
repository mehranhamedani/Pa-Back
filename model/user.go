package model

// User struct
type User struct {
	UserId    string `bson:"_id"`
	FirstName string `bson:"firstName"`
	LastName  string `bson:"lastName"`
	Email     string `bson:"email"`
	UserName  string `bson:"userName"`
	Password  string `bson:"password"`
}
