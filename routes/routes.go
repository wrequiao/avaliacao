package routes

import (
	"avaliacao/controller"
	"net/http"
)

func SetupRoutes() {
	http.HandleFunc("/", controller.Index)
	http.HandleFunc("/uploadfile", controller.UploadFile)
	http.HandleFunc("/upload", controller.Upload)
	http.HandleFunc("/processar", controller.Processar)
	http.HandleFunc("/delete", controller.Delete)
}
