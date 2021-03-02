package controller

import (
	"avaliacao/commons"
	"avaliacao/models"
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func UploadFile(w http.ResponseWriter, r *http.Request) {
	commons.GetTemplates().ExecuteTemplate(w, "Uploadfile", nil)
}

func Upload(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	var arq string = handler.Filename

	tempFile, err := ioutil.TempFile("temp", arq)
	if err != nil {
		fmt.Println(err)
	}

	defer tempFile.Close()
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	tempFile.Write(fileBytes)
	readFiles(tempFile.Name())
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func readFiles(arq string) {

	arquivoID := models.InsertArquivoGetId(arq)

	file, err := os.Open(arq)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		models.InsereArquivoLinha(arquivoID, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	e := os.Remove(arq)
	if e != nil {
		log.Fatal(e)
	}

}
