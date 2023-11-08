package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"path/filepath"
	"siiliboard/internal/marshal"

	"github.com/julienschmidt/httprouter"
)

var DEBUG bool = false

func NewRouter(rootPath string) *httprouter.Router {

	iPath := filepath.Join(rootPath, "static", "images") + "/"
	router := httprouter.New()
	router.ServeFiles("/images/*filepath", http.Dir(iPath))

	registerViews(router)
	return router
}

func registerViews(router *httprouter.Router) {

	router.GET("/", Index)
	router.GET("/boards", GetBoardsHandler)
	router.GET("/boards/:id", GetBoardHandler)
	router.POST("/boards", CreateBoardHandler)

	if DEBUG {
		router.DELETE("/boards", RemoveAllBoardsHandler)
	}
}

func readBodyToModel(r *http.Request, rt marshal.RequestTemplate) error {

	body, err := io.ReadAll(r.Body)

	if err != nil {
		log.Printf("Error reading request body: - %s\n", err.Error())
		return err
	}

	err = json.Unmarshal(body, rt)
	if err != nil {
		log.Printf("Error unmarshalling request data: - %s\n", err.Error())
		return err
	}

	err = rt.Validate()
	if err != nil {
		log.Printf("Error during request data validation - %s", err.Error())
		return err
	}

	return nil
}
