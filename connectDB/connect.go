package connectdb

import "database/sql"

func setupDB(dbDriver string, dsn string) (*sql.DB, error) {
    db, err := sql.Open(dbDriver, dsn)
    if err != nil {
        return nil, err
    }
    return db, err
}