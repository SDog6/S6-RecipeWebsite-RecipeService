package dbaccess

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func ConnectToDb() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("failed to load env", err)
	}

	db, err := sql.Open("mysql", os.Getenv("DSN"))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatalf("failed to ping: %v", err)
	}
	log.Println("Successfully connected to PlanetScale!")
	return db
}
