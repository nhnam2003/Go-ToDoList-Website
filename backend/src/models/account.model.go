package models

type Account struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
	Name     string `json:"name" bson:"name"`
	Age      int    `json:"age" bson:"age"`
	Role     string `json:"role" bson:"role"`
}
