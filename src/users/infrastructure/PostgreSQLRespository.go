package infrastructure

import (
	"fmt"
	"time"

	log "github.com/gofiber/fiber/v2/log"
	"github.com/jmoiron/sqlx"
	"github.com/luizarnoldch/mesapara2_backend/config/env"
)

// GetPostgreSQLClient retrieves a PostgreSQL client.
//
// It loads the configuration from the .env file and establishes a connection with the PostgreSQL database.
// It sets the connection properties and returns the PostgreSQL client.
//
// Returns:
// - *sqlx.DB: The PostgreSQL client.
func GetPostgreSQLClient() *sqlx.DB {
	config, err := env.LoadConfig(".env")
	if err != nil {
		log.Fatalf("error loading .env: %v", err)
	}

	dbUser := config.MICRO.DB.PSQL.USER
	dbPasswd := config.MICRO.DB.PSQL.PASS
	dbHost := config.MICRO.DB.PSQL.HOST
	dbPort := config.MICRO.DB.PSQL.PORT
	dbSchema := config.MICRO.DB.PSQL.SCHEMA

	dataSource := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPasswd, dbSchema)

	client, err := sqlx.Open("postgres", dataSource)
	if err != nil {
		log.Fatalf("error open connection with psql: %v",err)
	}

	err = client.Ping()
	if err != nil {
		log.Fatalf("error while ping psql connection: %v",err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	log.Info("Database successfully connected!")
	return client
}