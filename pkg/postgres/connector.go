package postgres

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func getDatabaseURL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_DATABASE"))
}

func GetConnection() (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), getDatabaseURL())

	if err != nil {
		log.Fatal("Error Connecting To The Database")
		return nil, err
	}

	return conn, nil
}
