package connect

import (
	"database/sql"
)

func SetupDB(dbDriver string, dsn string) (*sql.DB, error) {
    db, err := sql.Open(dbDriver, dsn)
    if err != nil {
        return nil, err
    }
    return db, err
}