
package database

import (
    "database/sql"
    "os"

    _ "github.com/mattn/go-sqlite3"
)

type GateSQLiteDatabase struct {
   DataFile string 
   Connection *sql.DB
}

func NewGateSQLiteDatabase (datafile string, purgeold bool) (*GateSQLiteDatabase, error) {
    if purgeold {
        os.Remove(datafile)
    }

    sql, err := sql.Open("sqlite3", datafile)

    if err != nil {
        return &GateSQLiteDatabase{}, err 
    }

    return &GateSQLiteDatabase{
        DataFile: datafile, 
        Connection: sql,
    }, nil 
}

func (gdb *GateSQLiteDatabase) BuildStructure () (bool) {
    return false
}

func (gdb *GateSQLiteDatabase) AddUser (userid string) (bool) {
    return false
}

func (gdb *GateSQLiteDatabase) RemoveUser (userid string) (bool) {
    return false 
}