package database

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
)

var DB *pgx.Conn

func Connect() {

	connStr := os.Getenv("DATABASE_URL")

	if connStr == "" {
		log.Fatal("DATABASE_URL not set")
	}

	var err error

	for i := 0; i < 15; i++ {

		DB, err = pgx.Connect(context.Background(), connStr)

		if err == nil {
			log.Println("Database connected successfully")
			return
		}

		log.Println("Waiting for database...")
		time.Sleep(3 * time.Second)
	}

	log.Fatal("Database connection failed:", err)
}