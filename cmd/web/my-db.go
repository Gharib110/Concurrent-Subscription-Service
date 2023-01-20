package main

import (
	"database/sql"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	"log"
	"os"
	"time"
)

func initDB() *sql.DB {
	conn := connectDB()
	if conn == nil {
		log.Panicln("Can not connect to postgresql :(...")
	}

	return conn
}

func connectDB() *sql.DB {
	counts := 5
	dsn := os.Getenv("postgre-dsn")

	for i := 0; i < counts; i++ {
		connection, err := openDB(dsn)
		if err != nil {
			log.Println("Postgresql is not ready yet :( ...")
		} else {
			log.Println("Postgresql is connected ;) ...")
			return connection
		}

		time.Sleep(time.Second * 3)
	}

	return nil
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, err
}
