package main

import (
	"log"

	"github.com/albiosz/honeycombs/internal/database"
	"github.com/albiosz/honeycombs/internal/util"
)

func main() {
	projectRoot := util.ProjectRoot()

	err := util.SetupEnvVariables(projectRoot + "/.env")
	if err != nil {
		log.Panicln(err)
	}

	db := database.Get()
	defer db.SqlDB.Close()
}
