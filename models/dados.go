package models

import (
	"avaliacao/db"
	"avaliacao/validacoes"
	"strings"
)

const (
	Cpf                = "Cpf"
	Private            = "Private"
	Incompleto         = "Incompleto"
	DataUltimaCompra   = "DataUltimaCompra"
	TicketMedio        = "TicketMedio"
	TicketUltimaCompra = "TicketUltimaCompra"
	LojaMaisFrequente  = "LojaMaisFrequente"
	LojaDaUltimaCompra = "LojaDaUltimaCompra"
)

type Dados struct {
	Id                 string //int
	Cpf                string
	Private            string //int
	Incompleto         string //int
	DataUltimaCompra   string //time.Time
	TicketMedio        string //float64
	TicketUltimaCompra string //float64
	LojaMaisFrequente  string
	LojaDaUltimaCompra string
	StatusValidacao    string
}

func GetDados(arquivoLinha []ArquivoLinha, idArquivo string) []Dados {

	db := db.ConectaComBancoDeDados()
	defer db.Close()
	query := `INSERT INTO dados (cpf, private, incompleto, dataultimacompra, ticketmedio, ticketultimacompra, lojamaisfrequente, lojadaultimacompra, idarquivo) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	dados := Dados{}
	dadosCollection := []Dados{}

	for index, element := range arquivoLinha {
		if index == 0 {
			continue
		}

		//fmt.Println(index)
		//fmt.Println(element.Linha)

		dados.Cpf = validacoes.Sanitize(getColumn(element.Linha, Cpf))
		dados.Private = getColumn(element.Linha, Private)
		dados.Incompleto = getColumn(element.Linha, Incompleto)
		dados.DataUltimaCompra = getColumn(element.Linha, DataUltimaCompra)
		dados.TicketMedio = getColumn(element.Linha, TicketMedio)
		dados.TicketUltimaCompra = getColumn(element.Linha, TicketUltimaCompra)
		dados.LojaMaisFrequente = validacoes.Sanitize(getColumn(element.Linha, LojaMaisFrequente))
		dados.LojaDaUltimaCompra = validacoes.Sanitize(getColumn(element.Linha, LojaDaUltimaCompra))
		dados.StatusValidacao = ""

		if !validacoes.IsCPF(dados.Cpf) {
			dados.StatusValidacao = "Cpf inválido/"
		}

		if !validacoes.IsCNPJ(dados.LojaDaUltimaCompra) {
			dados.StatusValidacao = dados.StatusValidacao + "Cnpj LojaDaUltimaCompra inválidos/"
		}

		if !validacoes.IsCNPJ(dados.LojaMaisFrequente) {
			dados.StatusValidacao = dados.StatusValidacao + "Cnpj LojaMaisFrequente inválidos/"
		}

		if dados.StatusValidacao == "" {
			dados.StatusValidacao = "Dados OK"
		}

		dadosCollection = append(dadosCollection, dados)

		query = `INSERT INTO dados (cpf, private, incompleto, dataultimacompra, ticketmedio, ticketultimacompra, lojamaisfrequente, lojadaultimacompra, status, idarquivo) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`
		stmt, err := db.Prepare(query)
		validacoes.CheckError(err)
		stmt.Exec(dados.Cpf, dados.Private, dados.Incompleto, dados.DataUltimaCompra, dados.TicketMedio, dados.TicketUltimaCompra, dados.LojaMaisFrequente, dados.LojaDaUltimaCompra, dados.StatusValidacao, idArquivo)

		defer stmt.Close()
		validacoes.CheckError(err)
	}

	return dadosCollection

}

func getColumn(linha, coluna string) string {
	retorno := ""
	switch coluna {
	case "Cpf":
		retorno = linha[0:18]
	case "Private":
		retorno = linha[19:30]
	case "Incompleto":
		retorno = linha[31:42]
	case "DataUltimaCompra":
		if strings.TrimSpace(linha[43:64]) == "NULL" {
			retorno = ""
		} else {
			retorno = linha[43:64]
		}
	case "TicketMedio":
		if strings.TrimSpace(linha[65:86]) == "NULL" {
			retorno = "0.0"
		} else {
			retorno = strings.ReplaceAll(linha[65:86], ",", ".")
		}
	case "TicketUltimaCompra":
		if strings.TrimSpace(linha[87:110]) == "NULL" {
			retorno = "0.0"
		} else {
			retorno = strings.ReplaceAll(linha[87:110], ",", ".")
		}
	case "LojaMaisFrequente":
		if strings.TrimSpace(linha[111:130]) == "NULL" {
			retorno = ""
		} else {
			retorno = linha[111:130]
		}
	case "LojaDaUltimaCompra":
		if strings.TrimSpace(linha[131:135]) == "NULL" {
			retorno = ""
		} else {
			retorno = linha[131:149]
		}
	default:
		retorno = ""
	}

	return strings.TrimSpace(retorno)
}
