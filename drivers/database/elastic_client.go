package database

import (
	"crud/config"
	"fmt"
	"log"

	"github.com/elastic/go-elasticsearch"
)

// ElasticDBClient the specific client for mongo db
type ElasticDBClient struct {
	connection *elasticsearch.Client
}

func (client *ElasticDBClient) connect(config config.Database) {
	// Initialize a client with the default settings.
	//
	// An `ELASTICSEARCH_URL` environment variable will be used when exported.
	//
	cfg := elasticsearch.Config{
		Addresses: []string{
			fmt.Sprintf("http://%v:%v", config.Host, config.Port),
		},
	}
	fmt.Printf("connection parameters %v", cfg)
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating Elastisearch client: %s", err)
	}

	client.connection = es
}

// GetConnection retrieve database connection as singleton
func (client *ElasticDBClient) GetConnection() *elasticsearch.Client {
	if !client.isConnected() {
		client.connect(config.Config.Elastic)
		fmt.Printf("Elastic db connected %v ", client.connection)
	}
	return client.connection
}

// isConnected return if there a database connection in memory
func (client ElasticDBClient) isConnected() bool {
	return client.connection != nil
}
