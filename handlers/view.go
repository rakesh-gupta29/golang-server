package handlers

import "net/http"

func HandleHome(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("this is home route"))
}

func HandleAbout(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("about us page will be rendered here"))
}
