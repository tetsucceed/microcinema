package dbo

import (
	"context"
	"fmt"
	"log"
	"model"
	"os"
	"time"

	"github.com/golang-migrate/migrate"
	_ "github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/jackc/pgx"
)

var DB *pgx.Conn
var DBUrl string

func GetConnectionDb() {
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")

	var ConnectionUrl = "postgres://%s:%s@%s/%s"
	FullUrl := fmt.Sprintf(ConnectionUrl, dbUser,
		dbPassword, dbHost, dbName)
	log.Println(FullUrl)

	conn, err := pgx.Connect(context.Background(), FullUrl)
	if err != nil {
		panic("failed to connect database")
	}

	DBUrl = FullUrl
	DB = conn
}

func ApplyMigration(dburl string) {
	m, err := migrate.New(
		"file://resources/migrations", dburl)
	if err != nil && err != migrate.ErrNoChange {
		log.Panic(err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Panic(err)
	}
	m.Steps(2)
}

func LoadMovies() []model.Movie {
	rows, err := DB.Query(context.Background(), "select username,is_on_duty,update_date from duty")
	var Items []model.Movie

	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		var value []interface{}
		value, err = rows.Values()
		Items = append(Items, model.Movie{
			ID:          value[0].(int),
			Name:        fmt.Sprintf("%v", value[1]),
			Poster:      fmt.Sprintf("%v", value[2]),
			MovieUrl:    fmt.Sprintf("%v", value[3]),
			IsPaid:      value[4].(bool),
			ReleaseDate: value[5].(time.Time),
			Genre:       fmt.Sprintf("%v", value[6])})
	}

	fmt.Println(Items)
	return Items
}
