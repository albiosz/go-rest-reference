package main

import (
	"log"

	"github.com/albiosz/honeycombs/internal/config"
	"github.com/albiosz/honeycombs/internal/database"
	"github.com/albiosz/honeycombs/internal/util"
)

func main() {
	projectRoot := util.ProjectRoot()

	err := config.SetupEnvVariables(projectRoot + "/.env")
	if err != nil {
		log.Panicln(err)
	}

	db := database.Get()
	defer db.SqlDB.Close()
}
