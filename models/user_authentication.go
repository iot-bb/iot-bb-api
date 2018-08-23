package models

import (
	"time"
  "gopkg.in/mgo.v2/bson"
)

// UserAuthentication for login
type UserAuthentication struct {
  ID          bson.ObjectId `bson:"_id" json:"id"`

  UserID      bson.ObjectId `bson:"user_id" json:"user_id"`
  AccessToken string        `bson:"access_token" json:"access_token"`
  ExpiredDate time.Time     `bson:"expired_date" json:"expired_date"`

  CreatedAt   time.Time     `bson:"created_date" json:"created_date"`
}
