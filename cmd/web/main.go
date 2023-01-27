package main

import (
	"fmt"
	"github.com/Gharib110/Concurrent-Subscription-Service/handlers"
	"log"
	"net/http"
	"os"
	"sync"
)

const WEBPORT = 85

func main() {
	// Initialization of Postgresql DB
	db := initDB()
	err := db.Ping()
	if err != nil {
		panic(err)
		return
	}

	// Initialize redis and session
	session := initSession()

	infoLogger := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errLogger := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Lshortfile)

	wg := sync.WaitGroup{}

	app := handlers.Config{
		Session: session,
		DB:      db,
		InfoLog: infoLogger,
		ErrLog:  errLogger,
		Wait:    &wg,
	}

	serve(&app)
}

func serve(app *handlers.Config) {
	srv := &http.Server{
		Addr:              fmt.Sprintf(":%s", WEBPORT),
		Handler:           app.Routes(),
		TLSConfig:         nil,
		ReadTimeout:       21,
		ReadHeaderTimeout: 0,
		WriteTimeout:      20,
		IdleTimeout:       20,
	}

	app.InfoLog.Println("Web server is started on " + string(WEBPORT))
	err := srv.ListenAndServe()
	if err != nil {
		app.ErrLog.Fatal(err)
		return
	}
}
