package main

import (
	"fmt"
	"main/controllers"
	"net/http"
)

func main() {
	fmt.Println("Server is running on port 8080")
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/result", controllers.Result)
	http.HandleFunc("/redirectindex", controllers.RedirectIndex)
	http.HandleFunc("/download", controllers.Download)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":8080", nil)

}
