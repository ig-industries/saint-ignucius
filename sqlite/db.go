package sqlite

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	// sqlite driver
	_ "github.com/mattn/go-sqlite3"
)

// NewDB creates a new SQLite 3 database connection from a path to the sql file.
func NewDB(DBPath string) (*sqlx.DB, error) {
	db, err := sqlx.Open("sqlite", DBPath)
	if err != nil {
		return nil, fmt.Errorf("[sqlite] %s", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("[sqlite] %s", err)
	}

	return db, nil
}
