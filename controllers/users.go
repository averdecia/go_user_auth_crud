package controllers

import (
	"crud/database/mongodb"
	"crud/entities"
	"crud/middleware"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/gin-gonic/gin"
)

// UsersController is the users controller class
type UsersController struct {
	RestController
}

// Auth function to authenticate a user
func (controller UsersController) Auth(c *gin.Context) {
	var userE entities.User
	var auth entities.Auth

	err := c.BindJSON(&auth)
	user, err := mongodb.GetOneBy(controller.MongoDBClient, Collections["users"],
		bson.M{"email": auth.Email, "password": auth.Password},
		userE)

	if err != nil {
		fmt.Printf("error %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "INVALID_CREDENTIALS"})
		return
	}

	token, _ := middleware.SetSessionToken(user.(*entities.User), c)
	user.(*entities.User).Tokens = token
	c.JSON(http.StatusOK, gin.H{"user": user})
}

// List function to find users from database
func (controller UsersController) List(c *gin.Context) {
	var user entities.User
	result, err := mongodb.GetAll(controller.MongoDBClient, Collections["users"], bson.M{}, user)
	if err != nil {
		fmt.Printf("error %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"users": result})
}

// GetByID function to find a specific user from database
func (controller UsersController) GetByID(c *gin.Context) {

	var userE entities.User
	user, err := mongodb.GetByID(controller.MongoDBClient, Collections["users"], c.Params.ByName("id"), userE)
	if err != nil {
		fmt.Printf("error %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Parameters"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}

// Create function to create a new user on database
func (controller UsersController) Create(c *gin.Context) {
	var userE entities.User
	err := c.BindJSON(&userE)

	if err != nil {
		fmt.Printf("error %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Parameters"})
		return
	}

	user, err := mongodb.Create(controller.MongoDBClient, Collections["users"], &userE)
	if err != nil {
		fmt.Printf("error %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Unable to save user data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

// Update function to update a specific user from database
func (controller UsersController) Update(c *gin.Context) {
	id := c.Params.ByName("id")
	objectID, _ := primitive.ObjectIDFromHex(id)

	var user entities.User
	err := c.BindJSON(&user)
	if err != nil {
		fmt.Printf("error %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Parameters"})
		return
	}

	var setElements bson.D
	if len(user.Firstname) > 0 {
		setElements = append(setElements, bson.E{Key: "firstname", Value: user.Firstname})
	}
	if len(user.Lastname) > 0 {
		setElements = append(setElements, bson.E{Key: "lastname", Value: user.Lastname})
	}
	if len(user.Email) > 0 {
		setElements = append(setElements, bson.E{Key: "email", Value: user.Email})
	}
	if len(user.Password) > 0 {
		setElements = append(setElements, bson.E{Key: "password", Value: user.Password})
	}
	if user.Phone != 0 {
		setElements = append(setElements, bson.E{Key: "phone", Value: user.Phone})
	}
	if user.Social != nil {
		setElements = append(setElements, bson.E{Key: "social", Value: user.Social})
	}
	if len(user.Tokens) > 0 {
		setElements = append(setElements, bson.E{Key: "tokens", Value: user.Tokens})
	}

	updatedUser, err := mongodb.UpdateOne(controller.MongoDBClient, Collections["users"],
		bson.M{"_id": objectID}, setElements, user)

	c.JSON(http.StatusOK, gin.H{"user": updatedUser})
}

// Delete function to remove a specific user from database
func (controller UsersController) Delete(c *gin.Context) {
	_, err := mongodb.DeleteByID(controller.MongoDBClient, Collections["users"], c.Params.ByName("id"))

	if err != nil {
		fmt.Printf("error %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Unable to remove object"})
		return
	}

	c.Status(http.StatusOK)
}
