package storage

import (
	"fmt"
	"log"
	"os"

	"github.com/elias-gill/poliapi/types"
	"github.com/jmoiron/sqlx"
)

type postgresStorage struct {
	conn  *sqlx.DB
	table string
}

// returns a new connection to postgres
func NewPostgreStorage() *postgresStorage {
	return &postgresStorage{
		conn:  connectToPostgres(),
		table: "usuarios",
	}
}

// get the user data of a given id
func (s *postgresStorage) GetById(user string) (*types.User, error) {
	var data types.User
	err := s.conn.Select(
		&data,
		"SELECT u.password, u.name FROM $1 u (WHERE user = $2)",
		s.table, user,
	)

	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (s *postgresStorage) Delete(user string) error {
	_, err := s.conn.DB.Query(
		"delete from $1 where nombre = $2 cascade",
		s.table, user,
	)
	return err
}

func (s *postgresStorage) Update(section string, user string, data any) error {
	_, err := s.conn.Query(
		"update $1 u set $2 = $3 where u.nombre = $4",
		s.table, section, data, user,
	)
	return err
}

func (s *postgresStorage) Insert(u types.User) error {
	_, err := s.conn.NamedExec(
		"insert into "+s.table+" (nombre, password, email) values (:nombre, :password, :email)",
		u,
	)
	if err != nil {
		return err
	}
	return nil
}

func (db *postgresStorage) userNameIsNotRepeated(name string) error {
	var count int
	// buscar en la db
	err := db.conn.QueryRow(
		"select count(*) from usuarios where nombre = $1",
		name).Scan(&count)

	if err != nil {
		return err
	}

	if count != 0 {
		return fmt.Errorf("Username is already in use")
	}
	return nil
}

// returns a new connection to the PostgreSQL database
func connectToPostgres() *sqlx.DB {
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
