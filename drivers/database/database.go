package database

// InitMongoDBCLient ...
func InitMongoDBCLient() *MongoDBClient {
	mongoClient := MongoDBClient{}
	mongoClient.GetConnection()
	return &mongoClient
}

// InitElasticDBCLient ...
func InitElasticDBCLient() *ElasticDBClient {
	elasticClient := ElasticDBClient{}
	elasticClient.GetConnection()
	return &elasticClient
}
