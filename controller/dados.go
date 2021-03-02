package controller

import (
	"avaliacao/commons"
	"avaliacao/models"
	"net/http"
)

func Processar(w http.ResponseWriter, r *http.Request) {
	var idArquivo = r.URL.Query().Get("id")
	arquivoLinha := models.GetArquivoLinha(idArquivo)
	dadosCollection := models.GetDados(arquivoLinha, idArquivo)
	models.AtualizaDataProcessamento(idArquivo)
	commons.GetTemplates().ExecuteTemplate(w, "Relatorio", dadosCollection)
}
