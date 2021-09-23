package controllers

import (
	"context"
	"crud/entities"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// ActivityController is the users controller class
type ActivityController struct {
	RestController
}

// List function to find users from database
func (controller ActivityController) List(c *gin.Context) {
	res, err := controller.ElasticDBClient.GetConnection().API.Info()
	if err != nil {
		fmt.Printf("conexion %v", err)
		c.Status(http.StatusBadRequest)
		return
	}

	fmt.Printf("conexion success %v", res)
	c.Status(http.StatusOK)
}

// GetByID function to find a specific user from database
func (controller ActivityController) GetByID(c *gin.Context) {
	id := c.Params.ByName("id")

	// Build the request body.
	query := strings.NewReader(`{
		"query": {
			"bool": {
				"must":[
					{
						"match": {
							"user": "` + id + `"
						}
					}
				],
				"filter":[
					{
						"range": {
							"timestamp": {
								"gte": 20
							}
						}
					}
				]
			}
		}
	}`)

	// Generate request
	api := controller.ElasticDBClient.GetConnection().API
	res, err := api.Search(
		api.Search.WithContext(context.Background()),
		api.Search.WithIndex("activity"),
		api.Search.WithBody(query),
		api.Search.WithTrackTotalHits(true),
		api.Search.WithPretty(),
	)
	// Validate response
	if err != nil {
		fmt.Printf("Error getting response: %s", err)
		c.Status(http.StatusBadRequest)
		return
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			fmt.Printf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and error information.
			fmt.Printf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
			c.Status(http.StatusBadRequest)
			return
		}
	}

	// Success case
	var result entities.ElasticSearchResults
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	var activities []entities.Activity
	for _, hit := range result.Hits.Hits {
		var activity entities.Activity
		err := json.Unmarshal(hit.Source, &activity)
		if err != nil {
			fmt.Printf("Error converting source response %v", err)
			c.Status(http.StatusBadRequest)
			return
		}
		activities = append(activities, activity)
	}

	c.JSON(http.StatusOK, gin.H{"activities": activities})
}

// Create function to create a new user on database
func (controller ActivityController) Create(c *gin.Context) {

}

// Update function to update a specific user from database
func (controller ActivityController) Update(c *gin.Context) {

}

// Delete function to remove a specific user from database
func (controller ActivityController) Delete(c *gin.Context) {

}
