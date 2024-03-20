package main

import (
	"main/controllers"
	"net/http"
)

func main() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/result", controllers.Result)
	http.HandleFunc("/redirectindex", controllers.RedirectIndex)
	http.HandleFunc("/download", controllers.Download)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":8080", nil)

}
