package db

import (
	"avaliacao/config"
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/lib/pq"
)

func ConectaComBancoDeDados() *sql.DB {
	var configuration = config.GetConfig() //define o modo de operacao, dev ou prod
	//configuration = GetConfig("prod")
	port, err := strconv.Atoi(configuration.DB_PORT)

	conexao := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", configuration.DB_HOST, port, configuration.DB_USERNAME, configuration.DB_PASSWORD, configuration.DB_NAME)
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}
