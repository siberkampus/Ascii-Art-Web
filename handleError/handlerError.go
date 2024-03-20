package handleerror

import "net/http"

func UndefinedUrl(w http.ResponseWriter, r *http.Request,path string)bool{
	if r.URL.Path!=path{
		http.NotFound(w,r)
		return true
	}
	return false
}

func StatusInternalServerError( w http.ResponseWriter,r *http.Request,err error){
	http.Error(w,err.Error(),http.StatusInternalServerError)
}
