package db

import (
	"database/sql"
	"embed"
	"log"
	"sync"

	"github.com/NotCoffee418/dbmigrator"
	"github.com/NotCoffee418/home-control-center/internal/config"

	_ "modernc.org/sqlite"
)

var (
	db   *sql.DB
	once sync.Once
)

//go:embed migrations/*.sql
var migrationFS embed.FS

// Initialize must be called manually on startup
func InitializeDatabase() {
	// Create DB before migrations
	db := GetDB()
	_, err := db.Exec("SELECT 1;")
	if err != nil {
		log.Printf("Warning: Could not create DB: %v", err)
	}

	// Apply migrations
	dbmigrator.SetDatabaseType(dbmigrator.SQLite)
	<-dbmigrator.MigrateUpCh(
		db,
		migrationFS,
		"migrations",
	)
}

func GetDB() *sql.DB {
	once.Do(func() {
		var err error
		db, err = sql.Open("sqlite", config.GetConfig().DatabasePath)
		if err != nil {
			log.Fatal(err)
		}
		// Verify connection
		if err = db.Ping(); err != nil {
			log.Fatal(err)
		}
	})
	return db
}
