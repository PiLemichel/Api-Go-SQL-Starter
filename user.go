package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
_ "github.com/jinzhu/gorm/dialects/mysql"
"golang.org/x/crypto/bcrypt"

)

type (
 // user describes a user type
 user struct {
  Id int `json:"id"`
  Login     string `json:"login"`
  Password string    `json:"password"`
 }
)

// createUser add a new user
func createUser(c *gin.Context) {
  password,_ := HashPassword(c.PostForm("password"))
	user := user{Login: c.PostForm("login"), Password: password}
	db.Save(&user)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "User item created successfully!"})
}

// fetchAllUser fetch all users
func fetchAllUser(c *gin.Context) {
	var users []user

	db.Find(&users)

	if len(users) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No user found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": users})
}

// fetchSingleUser fetch a single user
func fetchSingleUser(c *gin.Context) {
	var user user
	userID := c.Param("id")

	db.First(&user, userID)

	if user.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No user found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": user})
}

// updateUser update a user
func updateUser(c *gin.Context) {
	var user user
	userID := c.Param("id")

	db.First(&user, userID)

	if user.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No user found!"})
		return
	}

	db.Model(&user).Update("login", c.PostForm("login"))
  password,_ := HashPassword(c.PostForm("password"))
	db.Model(&user).Update("password", password)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "User updated successfully!"})
}

// deleteUser remove a user
func deleteUser(c *gin.Context) {
	var user user
	userID := c.Param("id")

	db.First(&user, userID)

	if user.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No user found!"})
		return
	}

	db.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "User deleted successfully!"})
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 5)
	return string(bytes), err
}
