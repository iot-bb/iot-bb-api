package v1

import (
  "encoding/hex"
  "net/http"
  "time"

  "golang.org/x/crypto/blake2b"
  "golang.org/x/crypto/bcrypt"

  "github.com/gin-gonic/gin"
  "github.com/iot-bb/iot-bb-api/dao"
  "github.com/iot-bb/iot-bb-api/models"

  "gopkg.in/mgo.v2/bson"
)

// Users Api
func Users(router *gin.RouterGroup) {
  db := dao.DAO()
  // Create
  router.POST("/users/create", func (c *gin.Context) {
    user := new(models.User)
    db.C("User").Find(
      bson.M{"email": c.PostForm("email")}).One(&user)
    if user.ID != "" {
      c.JSON(http.StatusOK, gin.H{"message": "User already exists.", "status": http.StatusBadRequest})
    } else {
      hash, err := bcrypt.GenerateFromPassword([]byte(c.PostForm("password")), bcrypt.MinCost)
      if err != nil {
          println(err)
      }
      userMap := models.User {
        ID: bson.NewObjectId(),
        FirstName: c.PostForm("first_name"),
        LastName: c.PostForm("last_name"),
        Email: c.PostForm("email"),
        Password: string(hash),
        UpdatedDate: time.Now(),
        CreatedDate: time.Now()}
      err1 := db.C("User").Insert(userMap)
      if err1 == nil {
        c.JSON(http.StatusOK, gin.H{"message": "Success", "status": http.StatusOK, "user_id": userMap.ID})
      } else {
        c.JSON(http.StatusOK, gin.H{"message": err.Error(), "status": http.StatusBadRequest})
      }
    }
  })

  // Update
  router.POST("/users/update", func (c *gin.Context) {
    // TODO: implement update user
    c.JSON(http.StatusOK, gin.H{"message": "Success", "status": http.StatusOK})
  })

  // Login by email and password
  router.POST("/users/login", func (c *gin.Context) {
    var users []*models.User
    db.C("User").Find(bson.M{
      "email": c.PostForm("email")}).All(&users)
    if len(users) > 0 {
      err := bcrypt.CompareHashAndPassword(
        []byte(users[0].Password),
        []byte(c.PostForm("password")))
      if err != nil {
        c.JSON(http.StatusOK, gin.H{"message": "Login failed.", "status": http.StatusBadRequest})
      } else {
        hash, _ := blake2b.New256([]byte(time.Now().String()))
        accessToken := hex.EncodeToString(hash.Sum(nil))
        var userAuthentications[] *models.UserAuthentication
        db.C("UserAuthentication").Find(
          bson.M{"user_id": users[0].ID}).All(&userAuthentications)
        for _, element := range userAuthentications {
          update := element
          update.ExpiredDate = time.Now()
          db.C("UserAuthentication").Update(bson.M{"_id": element.ID}, update)
        }
        userMap := models.UserAuthentication {
          ID: bson.NewObjectId(),
          UserID: users[0].ID,
          AccessToken: accessToken}
        db.C("UserAuthentication").Insert(userMap)
        c.JSON(http.StatusOK, gin.H{"message": "Success", "status": http.StatusOK, "access_token": accessToken})
      }
    } else {
      c.JSON(http.StatusBadRequest, gin.H{"message": "User doesn't found.", "status": http.StatusBadRequest})
    }
  })

  // Get User info by access token
  router.GET("/users/info", func (c *gin.Context) {
    var user = new(models.User)
    var userAuthentication models.UserAuthentication
    db.C("UserAuthentication").Find(
      bson.M{"access_token": c.GetHeader("Authorization")}).One(&userAuthentication)
    if userAuthentication.AccessToken != "" {
      db.C("User").FindId(userAuthentication.UserID).One(&user)
      c.JSON(http.StatusOK, gin.H{"message": "Success", "status": http.StatusOK, "user": user})
    } else {
      c.JSON(http.StatusOK, gin.H{"message": "No user found.", "status": http.StatusOK})
    }
  })
}
