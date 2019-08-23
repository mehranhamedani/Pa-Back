package model

// User struct
type User struct {
	FirstName string `bson:"firstName"`
	LastName  string `bson:"lastName"`
	Email     string `bson:"email"`
	UserName  string `bson:"userName"`
	Password  string `bson:"password"`
}
