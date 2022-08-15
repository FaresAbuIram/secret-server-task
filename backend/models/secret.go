package models

type Secret struct {
	Data   string `bson:"data"`
	Expire int    `bson:"expire" binding:"required"`
	Views  int    `bson:"views" binding:"required"`
}
