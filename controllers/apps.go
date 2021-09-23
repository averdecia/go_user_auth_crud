package controllers

import (
	"crud/database/mongodb"
	"crud/entities"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/gin-gonic/gin"
	"github.com/imdario/mergo"
)

// AppsController is the users controller class
type AppsController struct {
	RestController
}

// List function to find users from database
func (controller AppsController) List(c *gin.Context) {
	var config entities.App
	if user, ok := c.Get("Session"); ok {
		fmt.Printf("User %v", user)
	}
	result, err := mongodb.GetAll(controller.MongoDBClient, Collections["apps"], bson.M{}, config)
	if err != nil {
		fmt.Printf("error %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"apps": result})
}

// GetByID function to find a specific user from database
func (controller AppsController) GetByID(c *gin.Context) {

	var appE entities.App
	app, err := mongodb.GetByID(controller.MongoDBClient, Collections["apps"], c.Params.ByName("id"), appE)
	if err != nil {
		fmt.Printf("error %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Parameters"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"apps": app})
}

// GetBy function to find a specific user from database
func (controller AppsController) GetBy(c *gin.Context) {
	name := c.Params.ByName("fieldname")
	value := c.Params.ByName("fieldvalue")
	var bsonQuery bson.M
	if name == "id" {
		objectID, _ := primitive.ObjectIDFromHex(value)
		bsonQuery = bson.M{"_id": objectID}
	} else {
		bsonQuery = bson.M{name: value}
	}
	var config entities.App
	configs, err := mongodb.GetAll(controller.MongoDBClient, Collections["apps"], bsonQuery, config)
	if err != nil {
		fmt.Printf("error %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Parameters"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"apps": configs})
}

// Create function to create a new user on database
func (controller AppsController) Create(c *gin.Context) {
	var configE entities.App
	err := c.BindJSON(&configE)

	if err != nil {
		fmt.Printf("error %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Parameters"})
		return
	}

	config, err := mongodb.Create(controller.MongoDBClient, Collections["apps"], &configE)
	if err != nil {
		fmt.Printf("error %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Unable to save data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"apps": config})
}

// Update function to update a specific config from database
func (controller AppsController) Update(c *gin.Context) {
	id := c.Params.ByName("id")
	objectID, _ := primitive.ObjectIDFromHex(id)

	var app entities.App
	err := c.BindJSON(&app)
	if err != nil {
		fmt.Printf("error %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Parameters"})
		return
	}
	var setElements bson.D
	if len(app.Name) > 0 {
		setElements = append(setElements, bson.E{Key: "name", Value: app.Name})
	}
	if len(app.Icon) > 0 {
		setElements = append(setElements, bson.E{Key: "icon", Value: app.Icon})
	}
	if len(app.Package) > 0 {
		setElements = append(setElements, bson.E{Key: "package", Value: app.Package})
	}
	if len(app.BundleID) > 0 {
		setElements = append(setElements, bson.E{Key: "bundle_id", Value: app.BundleID})
	}
	if len(app.Description) > 0 {
		setElements = append(setElements, bson.E{Key: "description", Value: app.Description})
	}

	updatedConfig, err := mongodb.UpdateOne(controller.MongoDBClient, Collections["apps"],
		bson.M{"_id": objectID}, setElements, app)
	c.JSON(http.StatusOK, gin.H{"app": updatedConfig})
}

// Delete function to remove a specific user from database
func (controller AppsController) Delete(c *gin.Context) {
	_, err := mongodb.DeleteByID(controller.MongoDBClient, Collections["apps"], c.Params.ByName("id"))

	if err != nil {
		fmt.Printf("error %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Unable to remove object"})
		return
	}

	// Remove all versions related
	/*_, err := mongodb.DeleteAll(controller.MongoDBClient, Collections["versions"], bson.M{"app_id": c.Params.ByName("id")})
	if err != nil {
		fmt.Printf("error %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Unable to remove object"})
		return
	}*/

	c.Status(http.StatusOK)
}

// AddVersion function to remove a specific user from database
func (controller AppsController) AddVersion(c *gin.Context) {
	var versionE entities.Version
	err := c.BindJSON(&versionE)

	if err != nil {
		fmt.Printf("error %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Parameters"})
		return
	}

	var appE entities.App
	app, err := mongodb.GetByID(controller.MongoDBClient, Collections["apps"], versionE.AppID, appE)
	if err != nil {
		fmt.Printf("error %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Parameters"})
		return
	}

	var setElements bson.D
	for _, v := range app.(entities.App).Versions {
		if v.VersionString == versionE.VersionString {
			mergo.Merge(&versionE, v)
			setElements = append(setElements, bson.E{Key: "$set", Value: bson.M{"versions.$": versionE}})

			updatedConfig, err := mongodb.UpdateOne(controller.MongoDBClient, Collections["apps"],
				bson.M{"_id": app.(entities.App).ID, "versions.versionstring": v.VersionString},
				setElements, app)
			if err != nil {
				fmt.Printf("error %v", err)
				c.JSON(http.StatusBadRequest, gin.H{"message": "Unable to save on database"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"version": updatedConfig})
			return
		}
	}
	setElements = append(setElements, bson.E{Key: "$puch", Value: bson.M{"versions.$": versionE}})
	updatedConfig, err := mongodb.UpdateOne(controller.MongoDBClient, Collections["apps"],
		bson.M{"_id": app.(entities.App).ID}, setElements, app)

	if err != nil {
		fmt.Printf("error %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Unable to save on database"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"version": updatedConfig})

}

// RemoveVersion function to remove a specific app version
func (controller AppsController) RemoveVersion(c *gin.Context) {
	//app_id := c.Params.ByName("app_id")
	//versionstring := c.Params.ByName("versionstring")

}
