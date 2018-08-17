package dao

import (
  "os"
  "gopkg.in/mgo.v2"
)

// DAO database connection
func DAO() *mgo.Database {
  session, err := mgo.Dial(os.Getenv("MONGODB_URI"))
  print(err)
  print("===== Setup Database =====")
  print("- mongodb uri: " + os.Getenv("MONGODB_URI"))
  print("- mongodb name: " + os.Getenv("MONGODB_NAME"))
  return session.DB(os.Getenv("MONGODB_NAME"))
}

