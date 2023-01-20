package main

const WEBPORT = 85

func main() {
	// Initialization of Postgresql DB
	db := initDB()
	db.Ping()
}
