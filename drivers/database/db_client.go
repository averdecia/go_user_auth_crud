package database

// DBClient is the database connection
type DBClient struct {
	connection interface{}
}

func (client DBClient) connect() {
	// Todo: implement
}

// GetConnection retrieve database connection as singleton
func (client DBClient) GetConnection() interface{} {
	if !client.isConnected() {
		client.connect()
	}
	return client.connection
}

// isConnected return if there a database connection in memory
func (client DBClient) isConnected() bool {
	return client.connection != nil
}
