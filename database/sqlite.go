
package database

import (
    "database/sql"
    "os"

    _ "github.com/mattn/go-sqlite3"
)

func NewGateSQLiteDatabase (datafile string, purgeold bool) (*GateDatabase, error) {
    if purgeold {
        os.Remove(datafile)
    }

    sql, err := sql.Open("sqlite3", datafile)

    if err != nil {
        return &GateSQLiteDatabase{}, err 
    }

    return &GateDatabase{
        Database: datafile, 
        Conn: sql,
    }, nil 
}