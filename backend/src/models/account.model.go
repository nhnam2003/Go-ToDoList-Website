package models

type Account struct {
	username string `json:"username" bson:"username"`
	password string `json:"password" bson:"password"`
	name     string `json:"name" bson:"name"`
	age      int    `json:"age" bson:"age"`
	role     string `json:"role" bson:"role"`
}
