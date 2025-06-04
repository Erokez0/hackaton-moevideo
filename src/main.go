package main

import (
	"os"

	"github.com/Erokez0/hackaton-moevideo/src/categorizers/skydns"
	"github.com/Erokez0/hackaton-moevideo/src/config"
	"github.com/Erokez0/hackaton-moevideo/src/database"
	"github.com/Erokez0/hackaton-moevideo/src/server"
)

func main() {
	config.Init();
	database.Init();
	skydns.Init()
	
	args := os.Args[1:];
	if len(args) > 0 && args[0] == "seed" {
		database.Seed()
	}

	server.Run()
}