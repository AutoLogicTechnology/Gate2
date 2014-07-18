
package database

import (
	"database/sql"
)

type GateDatabase struct {
	Database *string 
	Conn *sql.DB 
}

func (gdb *GateDatabase) BuildStructure () (bool) {
    return false
}

func (gdb *GateDatabase) AddUser (userid string) (bool) {
    return false
}

func (gdb *GateDatabase) RemoveUser (userid string) (bool) {
    return false 
}