package dbaccess

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectToDb() *sql.DB {
	db, err := sql.Open("mysql", "wi1r1ff35xssfgt34lt7:pscale_pw_FZOQYetVkUt7hLRrt3ALIti7iUPmqcJfy1kHIJ3m08H@tcp(aws.connect.psdb.cloud)/recipedb?tls=true&parseTime=true")
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatalf("failed to ping: %v", err)
	}
	log.Println("Successfully connected to PlanetScale!")
	return db
}
