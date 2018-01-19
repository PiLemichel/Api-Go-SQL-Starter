package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
_ "github.com/jinzhu/gorm/dialects/mysql"
"golang.org/x/crypto/bcrypt"

)

func login(c *gin.Context) {
	var user user
	userID := c.Param("id")

	db.First(&user, userID)

	if user.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No user found!"})
		return
	}

  if(CheckPasswordHash(c.PostForm("password"),user.Password)){
    c.JSON(http.StatusNotFound, gin.H{"status": http.StatusOK, "message": "authorized"})
  }else{
    c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No authorized"})

  }
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
