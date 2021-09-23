package main

import (
	"crud/config"
	"crud/controllers"
	"crud/drivers/database"
	"crud/middleware"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	config.LoadConfig("config", "./src/crud/config")
	// database.Build(config.Database, config.DefaultDatabase)
	mongodb := database.InitMongoDBCLient()
	elasticdb := database.InitElasticDBCLient()

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"},
		AllowMethods:     []string{"PUT", "PATCH", "OPTIONS", "POST", "DELETE"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:4200"
		},
		MaxAge: 12 * time.Hour,
	}))
	usersR := router.Group("/users")
	{
		// Dependency injection
		controller := &controllers.UsersController{
			RestController: controllers.RestController{
				MongoDBClient: mongodb,
			},
		}

		usersR.POST("auth", controller.Auth)
		usersR.GET("", controller.List)
		usersR.GET("/:id", controller.GetByID)
		usersR.POST("", controller.Create)
		usersR.PUT("/:id", controller.Update)
		usersR.DELETE("/:id", controller.Delete)
	}
	activityR := router.Group("/activity")
	{
		// Dependency injection
		controller := &controllers.ActivityController{
			RestController: controllers.RestController{
				ElasticDBClient: elasticdb,
			},
		}

		activityR.GET("", controller.List)
		activityR.GET("/:id", controller.GetByID)
		activityR.POST("", controller.Create)
		activityR.PUT("/:id", controller.Update)
		activityR.DELETE("/:id", controller.Delete)
	}

	mobileConfigR := router.Group("/config").Use(middleware.TokenUserAuthorization())
	{
		// Dependency injection
		controller := &controllers.MobileConfigController{
			RestController: controllers.RestController{
				MongoDBClient: mongodb,
			},
		}

		mobileConfigR.GET("", controller.List)
		mobileConfigR.GET("/:fieldname/:fieldvalue", controller.GetBy)
		mobileConfigR.GET("/:fieldname", controller.GetByID)
		mobileConfigR.POST("", controller.Create)
		mobileConfigR.PUT("/:id", controller.Update)
		mobileConfigR.DELETE("/:id", controller.Delete)
	}

	appsR := router.Group("/apps").Use(middleware.TokenUserAuthorization())
	{
		// Dependency injection
		controller := &controllers.AppsController{
			RestController: controllers.RestController{
				MongoDBClient: mongodb,
			},
		}

		appsR.GET("", controller.List)
		appsR.GET("/:fieldname/:fieldvalue", controller.GetBy)
		appsR.GET("/:fieldname", controller.GetByID)
		appsR.POST("", controller.Create)
		appsR.PUT("/:id", controller.Update)
		appsR.DELETE("/:id", controller.Delete)
		appsR.POST("/version", controller.AddVersion)
	}
	router.Run()
}
