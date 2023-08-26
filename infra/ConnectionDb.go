package infra

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitDb() {
	//informações do banco
	username := "root"
	password := "admin"
	host := "localhost"
	port := "3306"
	dbName := "apidb"

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, dbName)

	var err error
	db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	//Verificar a conexão com o banco
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Conexão com o banco de dados estabelecida com sucesso.")
}

func CloseDB() {
	if db != nil {
		db.Close()
	}
}