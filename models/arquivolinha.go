package models

import (
	"avaliacao/db"
	"avaliacao/validacoes"
)

type ArquivoLinha struct {
	Id        int
	Linha     string
	IdArquivo int
}

func GetArquivoLinha(idArquivo string) []ArquivoLinha {

	db := db.ConectaComBancoDeDados()

	defer db.Close()

	query := `select id, linha, idarquivo from arquivolinha where idarquivo = $1 order by id`
	stmt, err := db.Prepare(query)
	validacoes.CheckError(err)
	selectArquivoLinha, err := stmt.Query(idArquivo)
	validacoes.CheckError(err)

	p := ArquivoLinha{}
	arquivoLinha := []ArquivoLinha{}

	for selectArquivoLinha.Next() {
		var id, IDArquivo int
		var linha string

		err = selectArquivoLinha.Scan(&id, &linha, &IDArquivo)
		validacoes.CheckError(err)

		p.Id = id
		p.Linha = linha
		p.IdArquivo = IDArquivo

		arquivoLinha = append(arquivoLinha, p)
	}

	return arquivoLinha
}

func InsereArquivoLinha(arquivoID int, linha string) {

	db := db.ConectaComBancoDeDados()
	defer db.Close()
	query := `insert into arquivolinha (linha, idarquivo) values($1, $2)`
	_, err := db.Exec(query, linha, arquivoID)
	validacoes.CheckError(err)

}
