package main

import (
	"library/db"
	"library/routes"
)

func main() {
	db.Init()
	e := routes.Init()
	e.Logger.Fatal(e.Start("localhost:3000"))
}
