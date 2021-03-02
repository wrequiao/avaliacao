package controller

import (
	"avaliacao/commons"
	"avaliacao/models"
	"net/http"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	var idArquivo = r.URL.Query().Get("id")
	models.DeletaArquivo(idArquivo)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func Index(w http.ResponseWriter, r *http.Request) {
	arquivos := models.SelectTodosOsArquivos()
	commons.GetTemplates().ExecuteTemplate(w, "Index", arquivos)
}
