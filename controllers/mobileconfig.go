package controllers

import (
	"crud/database/mongodb"
	"crud/entities"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/gin-gonic/gin"
)

// MobileConfigController is the users controller class
type MobileConfigController struct {
	RestController
}

// List function to find users from database
func (controller MobileConfigController) List(c *gin.Context) {
	var config entities.MobileConfig
	if user, ok := c.Get("Session"); ok {
		fmt.Printf("User %v", user)
	}
	result, err := mongodb.GetAll(controller.MongoDBClient, Collections["config"], bson.M{}, config)
	if err != nil {
		fmt.Printf("error %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"configs": result})
}

// GetByID function to find a specific user from database
func (controller MobileConfigController) GetByID(c *gin.Context) {

	var configE entities.MobileConfig
	config, err := mongodb.GetByID(controller.MongoDBClient, Collections["config"], c.Params.ByName("fieldname"), configE)
	if err != nil {
		fmt.Printf("error %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Parameters"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"config": config})
}

// GetBy function to find a specific user from database
func (controller MobileConfigController) GetBy(c *gin.Context) {
	name := c.Params.ByName("fieldname")
	value := c.Params.ByName("fieldvalue")
	var bsonQuery bson.M
	if name == "id" {
		objectID, _ := primitive.ObjectIDFromHex(value)
		bsonQuery = bson.M{"_id": objectID}
	} else {
		bsonQuery = bson.M{name: value}
	}
	var config entities.MobileConfig
	configs, err := mongodb.GetAll(controller.MongoDBClient, Collections["config"], bsonQuery, config)
	if err != nil {
		fmt.Printf("error %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Parameters"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"configs": configs})
}

// Create function to create a new user on database
func (controller MobileConfigController) Create(c *gin.Context) {
	var configE entities.MobileConfig
	err := c.BindJSON(&configE)

	if err != nil {
		fmt.Printf("error %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Parameters"})
		return
	}

	config, err := mongodb.Create(controller.MongoDBClient, Collections["config"], &configE)
	if err != nil {
		fmt.Printf("error %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Unable to save data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"config": config})
}

// Update function to update a specific config from database
func (controller MobileConfigController) Update(c *gin.Context) {
	id := c.Params.ByName("id")
	objectID, _ := primitive.ObjectIDFromHex(id)

	var config entities.MobileConfig
	err := c.BindJSON(&config)
	if err != nil {
		fmt.Printf("error %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Parameters"})
		return
	}
	/*
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

			updatedConfig, err := mongodb.UpdateOne(controller.MongoDBClient, Collections["config"],
				bson.M{"_id": objectID}, setElements, config)
		c.JSON(http.StatusOK, gin.H{"config": updatedConfig})
	*/

	c.JSON(http.StatusOK, gin.H{"id": objectID})
}

// Delete function to remove a specific user from database
func (controller MobileConfigController) Delete(c *gin.Context) {
	_, err := mongodb.DeleteByID(controller.MongoDBClient, Collections["config"], c.Params.ByName("id"))

	if err != nil {
		fmt.Printf("error %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Unable to remove object"})
		return
	}

	c.Status(http.StatusOK)
}
