package models

import (
	"time"
)

type Secret struct {
	ExpireDate time.Time `bson:"expire_date" binding:"required"`
	Views      int       `bson:"views" binding:"required"`
}

type Data struct {
	Data   string `bson:"data" binding:"required"`
	Views  int    `bson:"views" binding:"required"`
	Expire int    `bson:"expire" binding:"required"`
}

type ResponseData struct {
	Data    string `bson:"data"`
	Message string `bson:"message"`
	Object  Secret
}

type ResultToken struct {
	Data string `bson:"data"`
}
