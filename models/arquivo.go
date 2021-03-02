package models

import (
	"avaliacao/db"
	"avaliacao/validacoes"
)

type Arquivo struct {
	Id   int
	Nome string
}

func SelectTodosOsArquivos() []Arquivo {

	db := db.ConectaComBancoDeDados()

	defer db.Close()

	TodosOsArquivos, err := db.Query("select id, nome from arquivo where dataprocessamento is null order by id")
	validacoes.CheckError(err)

	p := Arquivo{}
	arquivos := []Arquivo{}

	for TodosOsArquivos.Next() {
		var id int
		var nome string

		err = TodosOsArquivos.Scan(&id, &nome)
		validacoes.CheckError(err)

		p.Id = id
		p.Nome = nome

		arquivos = append(arquivos, p)
	}

	return arquivos
}

func AtualizaDataProcessamento(idArquivo string) {

	db := db.ConectaComBancoDeDados()

	defer db.Close()

	query := `update arquivo set dataprocessamento=now() where id = $1`
	stmt, err := db.Prepare(query)
	validacoes.CheckError(err)
	stmt.Exec(idArquivo)
	defer stmt.Close()
	validacoes.CheckError(err)

	defer db.Close()

}

func DeletaArquivo(id string) {

	db := db.ConectaComBancoDeDados()

	defer db.Close()
	query := `delete from arquivolinha where idarquivo = $1`
	stmt, err := db.Prepare(query)
	validacoes.CheckError(err)
	stmt.Query(id)
	validacoes.CheckError(err)

	defer stmt.Close()

	query = `delete from dados where idarquivo = $1`
	stmt, err = db.Prepare(query)
	validacoes.CheckError(err)
	stmt.Query(id)
	validacoes.CheckError(err)

	defer stmt.Close()

	query = `delete from arquivo where id = $1`
	stmt, err = db.Prepare(query)
	validacoes.CheckError(err)
	stmt.Query(id)
	validacoes.CheckError(err)

	defer stmt.Close()

}

func InsertArquivoGetId(arq string) int {

	db := db.ConectaComBancoDeDados()
	defer db.Close()
	query := `insert into arquivo (nome, datacriacao) values($1, now()) RETURNING id`
	stmt, err := db.Prepare(query)
	validacoes.CheckError(err)
	defer stmt.Close()
	var arquivoID int = 0
	err = stmt.QueryRow(arq).Scan(&arquivoID)
	validacoes.CheckError(err)
	return arquivoID

}
