package main

import (
	"log"

	"github.com/albiosz/honeycombs/internal/database"
	"github.com/albiosz/honeycombs/internal/database/seed"
	"github.com/albiosz/honeycombs/internal/util"
)

func main() {
	err := util.SetupEnvVariables(util.ProjectRoot() + "/.env")
	if err != nil {
		log.Fatal(err)
	}

	db := database.Get()
	defer db.SqlDB.Close()

	db.Clear()
	seed.InsertAll(db.SqlDB)
}
