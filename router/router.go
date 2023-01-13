package router

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rakesh-gupta29/golang-server/handlers"
)

func handleHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("this is home route"))
}
func Mount() http.Handler {

	router := httprouter.New()

	router.NotFound = http.HandlerFunc(handlers.ErrorNotFound)
	router.MethodNotAllowed = http.HandlerFunc(handlers.ErrorMethodNotAllowed)

	router.ServeFiles("/static/*filepath", http.Dir("static"))
	router.HandlerFunc(http.MethodGet, "/", handlers.HandleHome)
	router.HandlerFunc(http.MethodGet, "/about", handlers.HandleAbout)

	return router
}
