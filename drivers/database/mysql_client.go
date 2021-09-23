package database

import (
	"crud/config"
	"database/sql"
	"fmt"
	"log"

	// mysql not needed to set variable, because is used inside database/sql
	_ "github.com/go-sql-driver/mysql"
)

// MysqlDBClient the specific client for mongo db
type MysqlDBClient struct {
	connection interface{}
}

func (client MysqlDBClient) connect(config config.Database) {
	fmt.Println("entro")
	fmt.Println(fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Database))
	db, err := sql.Open(config.Driver,
		fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
			config.User,
			config.Password,
			config.Host,
			config.Port,
			config.Database))
	if err != nil {
		log.Fatal(err)
	}
	// Test connection to Mysql
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	client.connection = db
	fmt.Println("Mysql db connected")
}

// GetConnection retrieve database connection as singleton
func (client MysqlDBClient) GetConnection(config config.Database) interface{} {
	if !client.isConnected() {
		client.connect(config)
	}
	return client.connection
}

// isConnected return if there a database connection in memory
func (client MysqlDBClient) isConnected() bool {
	return client.connection != nil
}
