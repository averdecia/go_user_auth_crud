package controllers

import (
	"crud/drivers/database"

	"github.com/gin-gonic/gin"
)

// Collections is used to access to database mongo collections
var Collections = map[string]string{
	"users":  "users",
	"config": "config",
	"apps":   "apps",
}

// RestController class
type RestController struct {
	MongoDBClient   *database.MongoDBClient
	ElasticDBClient *database.ElasticDBClient
}

// APIRestController rest controllers interface
type APIRestController interface {
	List(c *gin.Context)
	GetByID(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}
