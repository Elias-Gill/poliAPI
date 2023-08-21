package storage

import (
	"fmt"
	"log"
	"os"

	"github.com/elias-gill/poliapi/types"
	_ "github.com/lib/pq"
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

func (s *postgresStorage) Update(id string, user *types.User) error {
	query := fmt.Sprintf(
		`update %s u set 
        email=:email, nombre=:nombre, password=:password 
        where u.nombre = %s`,
		s.table, id,
	)

	_, err := s.conn.NamedExec(query, user)
	return err
}

func (s *postgresStorage) Insert(u *types.User) error {
	query := fmt.Sprintf(
		`insert into %s (nombre, password, email) 
        values (:nombre, :password, :email)`,
		s.table,
	)
	_, err := s.conn.NamedExec(query, u)
	return err
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
	db, err := sqlx.Open("postgres", pg_dsn)
	if err != nil {
		log.Fatal("No se pudo conectar la DB: ", err.Error())
	}

	//Call db.Ping() to check the connection
	pingErr := db.Ping()
	if err != nil {
		log.Fatal("No se pudo conectar con la DB: ", pingErr.Error())
	}

    log.Println("Conectado a DB postgres")
	return db
}
