package main

import (
	"flag"
	"log"
	"trash-stack-todo/model"
	"trash-stack-todo/routes"
)

func main() {
	migrate := flag.Bool(
		"migrate", false, "Crea las tablas en la base de datos",
	) // Parseando todas las flags
	flag.Parse()
	if *migrate {
		if err := model.MakeMigrations(); err != nil {
			log.Fatal(err)
		}
	}
	routes.SetupAndRun()
}
