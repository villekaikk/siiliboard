package routes

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func NewRouter() *httprouter.Router {

	router := httprouter.New()
	router.ServeFiles("/images/*filepath", http.Dir("./static/images/"))

    registerViews(router)
	return router
}

func registerViews(router *httprouter.Router) {

	router.GET("/", Index)
	
}