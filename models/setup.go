package models

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/glebarez/sqlite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

  var database *gorm.DB
  var err error

  dbType := os.Getenv("DB_TYPE")

  switch dbType {
  case "postgres":
    dsn := fmt.Sprintf("host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=America/Recife",
      os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))
    database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
  case "sqlite":

    database, err = gorm.Open(sqlite.Open(os.Getenv("SQLITE_DATABASE")), &gorm.Config{})
  default:
    slog.Error("Unsupported DB_TYPE. Use 'postgres' or 'sqlite'.", err)
  }

  if err != nil {
    panic("Failed to connect to database")
  }

  err = database.AutoMigrate(&ProjectDesc{})
  if err != nil {
    slog.Error("Error migrating", err)
  }
 // database.AutoMigrate(&ProjectQuota{})
 // database.AutoMigrate(&ProjectQuotaUsage{})
 // database.AutoMigrate(&ServerDesc{})
 // database.AutoMigrate(&ServerSpec{})
 // database.AutoMigrate(&ServerUsage{})
 // database.AutoMigrate(&ServerOwnership{})
 // database.AutoMigrate(&FlavorDesc{})
 // database.AutoMigrate(&FlavorSpec{})
 // database.AutoMigrate(&UserDesc{})
 // database.AutoMigrate(&UserProject{})

  DB = database
}
