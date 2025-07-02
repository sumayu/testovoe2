package bd

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func Database() (*sql.DB, error) {
	configPath := getConfigPath()
	
	if err := godotenv.Load(configPath); err != nil {
		log.Printf("Notice: config file not found at %s, using environment variables only", configPath)
	}

	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DATABASE")
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")

	if user == "" || password == "" || dbname == "" || host == "" || port == "" {
		return nil, fmt.Errorf("missing required database connection parameters")
	}

	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		user, password, dbname, host, port)

	log.Println("Connecting to database with DSN:", dsn)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	return db, nil
}

func getConfigPath() string {
	if os.Getenv("IS_DOCKER") == "TRUE" {
		return "/app/configs/config.env" 
	}
	return "./configs/config.env"      
}