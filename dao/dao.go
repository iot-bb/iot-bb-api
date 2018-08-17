// This program provides a sample application for using MongoDB with
// the mgo driver.
package dao

import (
  "gopkg.in/mgo.v2"
)

function DAO() MongoDB{
  session, err := mgo.Dial(os.Getenv("MONGODB_URI"))
  return session.DB(database)
}

