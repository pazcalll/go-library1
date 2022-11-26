package main

import "library/routes"

func main() {
	e := routes.Init()
	e.Logger.Fatal(e.Start("localhost:3000"))
}
