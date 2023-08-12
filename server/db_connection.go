package server

import (
	"os"

	//We are using the pgx driver to connect to PostgreSQL
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
)

var db = connectToDB()

// returns the current connection to the PostgreSQL database
func GetDBConnection() *sqlx.DB {
	return db
}

func connectToDB() *sqlx.DB {
	wd, _ := os.Getwd()
	godotenv.Load(wd + "/.env")
	pg_dsn := os.Getenv("PG_DSN")

	//Use sql.Open to initialize a new sql.DB object
	db, err := sqlx.Open("pgx", pg_dsn)
	if err != nil {
		log.Fatal("No se pudo conectar la DB: ", err.Error())
	}

	//Call db.Ping() to check the connection
	pingErr := db.Ping()
	if err != nil {
		log.Fatal("No se pudo conectar con la DB: ", pingErr.Error())
	}

	return db
}
