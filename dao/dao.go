// This program provides a sample application for using MongoDB with
// the mgo driver.
package dao

import (
  "os"
  "gopkg.in/mgo.v2"
)

// DAO database connection
func DAO() *mgo.Database {
  session, err := mgo.Dial(os.Getenv("MONGODB_URI"))
  print(err)
  return session.DB(os.Getenv("MONGODB_URI"))
}

