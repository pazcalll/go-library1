package main

import (
	"database/sql"
	"flag"
	"library/config"
	"library/db"
	"library/db/seeder"
	"library/routes"
	"log"
)

func main() {
	handleArgs()
}

func handleArgs() {
	flag.Parse()
	args := flag.Args()

	if len(args) >= 1 {
		switch args[0] {
		case "seed":
			conf := config.GetConfig()

			connectionString := conf.DB_USERNAME + ":" + conf.DB_PASSWORD + "@tcp(" + conf.DB_HOST + ":" + conf.DB_PORT + ")/" + conf.DB_NAME

			// connect DB
			db, err := sql.Open("mysql", connectionString)
			if err != nil {
				log.Fatalf("Error opening DB: %v", err)
			}
			seeder.Execute(db, args[1:]...)
		case "start":
			db.Init()
			e := routes.Init()
			e.Logger.Fatal(e.Start("localhost:3000"))
		}
	}
}
