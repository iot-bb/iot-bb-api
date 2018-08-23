package dao

import (
  "os"
  "gopkg.in/mgo.v2"
)

// DAO database connection
func DAO() *mgo.Database {
  // export MONGODB_URI=localhost:27017
  // export MONGODB_NAME=iotbb
  session, err := mgo.Dial(os.Getenv("MONGODB_URI"))
  if err != nil {
    print(err)
  }
  println("===== Setup Database =====")
  println("- mongodb uri: " + os.Getenv("MONGODB_URI"))
  println("- mongodb name: " + os.Getenv("MONGODB_NAME"))
  return session.DB(os.Getenv("MONGODB_NAME"))
}

