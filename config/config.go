package config

import (
	"database/sql"
	"os"
)
type Context struct {
	Db *sql.DB
	ErrorLogFile *os.File
}
