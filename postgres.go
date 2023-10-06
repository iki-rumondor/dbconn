package dbconn

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type postgresDB struct {
	db     *sql.DB
	gormDB *gorm.DB
}

var DB_HOST string
var DB_PORT string
var DB_USER string
var DB_PASS string
var DB_NAME string
var SSL_MODE string

func getPostgresStringConnection() (string, error) {
	if err := godotenv.Load(); err != nil {
		return "", err
	}

	DB_HOST = os.Getenv("PG_DB_HOST")
	DB_PORT = os.Getenv("PG_DB_PORT")
	DB_USER = os.Getenv("PG_DB_USER")
	DB_PASS = os.Getenv("PG_DB_PASSWORD")
	DB_NAME = os.Getenv("PG_DB_NAME")
	SSL_MODE = os.Getenv("PG_SSL_MODE")

	strConn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", DB_HOST, DB_PORT, DB_USER, DB_PASS, DB_NAME, SSL_MODE)
	return strConn, nil
}

func NewPostgresDB() (*sql.DB, error) {
	strcon, err := getPostgresStringConnection()
	if err != nil {
		return nil, err
	}

	db, err := sql.Open("postgres", strcon)
	if err != nil {
		return nil, err
	}

	defer db.Close()

	if err := db.Ping(); err != nil {
		return nil, err
	}

	fmt.Println("Successfully connected to postgres database")
	pg := postgresDB{
		db: db,
	}

	return pg.db, nil
}

func NewPostgresGormDB() (*gorm.DB, error) {
	strcon, err := getPostgresStringConnection()
	if err != nil {
		return nil, err
	}

	db, err := gorm.Open(postgres.Open(strcon), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	fmt.Println("Successfully connected to postgres database")
	pg := postgresDB{
		gormDB: db,
	}

	return pg.gormDB, nil
}
