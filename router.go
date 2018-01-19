package main

import (
       "github.com/gin-gonic/gin"
       "github.com/jinzhu/gorm"
       _ "github.com/jinzhu/gorm/dialects/mysql"
)



var db *gorm.DB

func init() {
 //open a db connection
 var err error
 db, err = gorm.Open("mysql", "root:root@tcp(localhost:8889)/user")
 if err != nil {
  panic("failed to connect database")
 }
}

func main() {
router := gin.Default()

user := router.Group("/user")
 {
  user.POST("/", createUser)
  user.GET("/", fetchAllUser)
  user.GET("/:id", fetchSingleUser)
  user.PUT("/:id", updateUser)
  user.DELETE("/:id", deleteUser)
 }

root := router.Group("/")
{
  root.POST("/login", login)
}

 router.Run(":3000")
}
