package models

import (
	"time"
  "gopkg.in/mgo.v2/bson"
)

// User for login
type User struct {
  ID          bson.ObjectId `bson:"_id" json:"id"`

  FirstName   string        `bson:"first_name" json:"first_name"`
  LastName    string        `bson:"last_name" json:"last_name"`
  Email       string        `bson:"email" json:"email"`
  Password    string        `bson:"password" json:"password"`

  UpdatedDate time.Time     `bson:"updated_date" json:"updated_date"`
  CreatedDate time.Time     `bson:"created_date" json:"created_date"`
}
