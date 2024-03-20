package controllers

import (
	"html/template"
	"main/calculate"
	handleerror "main/handleError"

	"net/http"
	"os"
	"strings"
)

func Index(w http.ResponseWriter, r *http.Request) {

	if handleerror.UndefinedUrl(w, r, "/") {
		return
	}
	view, err := template.ParseFiles("views/show/index.html", "views/show/content.html", "views/show/header.html","views/show/footer.html")
	if err != nil {
		handleerror.StatusInternalServerError(w, r, err)
		return
	}
	view.ExecuteTemplate(w, "index", nil)
}
func Result(w http.ResponseWriter, r *http.Request) {
	if handleerror.UndefinedUrl(w, r, "/result") {
		return
	}
	file_name := r.FormValue("file")
	arguman := r.FormValue("arguman")
	data := make(map[string][]interface{})

	checkbox := r.Form.Get("check")

	kelimeler := strings.Split(arguman, "\n") // harfleri aldık

	filenames := []string{"standard", "thinkertoy", "shadow"}
	if len(checkbox) == 0 {
		calculate.Yazdir(kelimeler, data, file_name)
	} else {
		for _, files := range filenames {
			calculate.Yazdir(kelimeler, data, files)
		}
	}
	view, err := template.ParseFiles("views/result/index.html", "views/result/header.html", "views/result/content.html","views/result/footer.html")
	if err != nil {
		handleerror.StatusInternalServerError(w, r, err)
		return
	}

	err = view.ExecuteTemplate(w, "index", data)
	if err != nil {
		handleerror.StatusInternalServerError(w, r, err)
		return
	}
}
func Download(w http.ResponseWriter, r *http.Request) {
	if handleerror.UndefinedUrl(w, r, "/download") {
		return
	}
	file, err := os.Create("result.txt")
	if err != nil {
		handleerror.StatusInternalServerError(w, r, err)
		return
	}
	defer file.Close()
	data := r.URL.Query().Get("data")
	file.Write([]byte(data))
	filepath := "result.txt"
	file_download, err := os.Open(filepath)
	if err != nil {
		handleerror.StatusInternalServerError(w, r, err)
		return
	}
	defer file_download.Close()
	fileinfo, err := file_download.Stat()
	if err != nil {
		handleerror.StatusInternalServerError(w, r, err)
		return
	}
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Content-Disposition", "attachment; filename="+filepath)
	http.ServeContent(w, r, filepath, fileinfo.ModTime(), file)

}
func RedirectIndex(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
