package test

import (
	"github.com/albiosz/honeycombs/internal/config"
	"github.com/albiosz/honeycombs/internal/database"
	"github.com/albiosz/honeycombs/internal/database/seed"
	"github.com/albiosz/honeycombs/internal/util"
)

func SetupTestDB() *database.DB {
	config.SetupEnvVariables(util.ProjectRoot() + "/.env")
	db := database.Get()
	db.Clear()
	seed.InsertAll(db.SqlDB)
	return db
}
