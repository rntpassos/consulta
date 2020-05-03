package repo

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "postgres"
	DB_NAME     = "consulta_db"
)

var Db *sql.DB

func AbreConexaoComBancoDeDadosSQL() (err error) {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)
	err = nil
	Db, err = sql.Open("postgres", dbinfo)
	if err != nil {
		return
	}

	err = Db.Ping()
	if err != nil {
		return
	}
	return
}
