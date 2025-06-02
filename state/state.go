package state

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type State struct {
	Env      Env
	Database *sql.DB
}

func (s *State) GetDatabaseURL() string {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		s.Env.DatabaseUser,
		s.Env.DatabasePassword,
		s.Env.DatabaseHost,
		strconv.Itoa(s.Env.DatabasePort),
		s.Env.DatabaseName)
	return dsn
}

func LoadState() State {
	var s State

	env := loadEnv()
	s.Env = env

	db, err := connectToDatabase(&s)
	if err != nil {
		log.Fatal("Something went wrong when connecting to the database")
	}

	s.Database = db

	return s
}
