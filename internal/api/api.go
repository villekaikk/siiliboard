package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"path/filepath"
	"siiliboard/internal/marshal"
	"siiliboard/internal/middleware"

	"github.com/gorilla/schema"
	"github.com/julienschmidt/httprouter"
)

var DEBUG bool = false
var decoder = schema.NewDecoder()

func NewRouter(rootPath string) *httprouter.Router {

	iPath := filepath.Join(rootPath, "static", "images") + "/"
	router := httprouter.New()
	router.ServeFiles("/images/*filepath", http.Dir(iPath))

	cssPath := filepath.Join(rootPath, "static", "styles") + "/"
	router.ServeFiles("/styles/*filepath", http.Dir(cssPath))

	registerViews(router)
	return router
}

func registerViews(router *httprouter.Router) {

	router.GET("/", Index)

	router.GET("/newboard", middleware.LogHTTP(GetNewBoardModal))
	router.GET("/boards", middleware.LogHTTP(GetBoardsHandler))
	router.GET("/boards/:bid", middleware.LogHTTP(GetBoardHandler))
	router.POST("/boards", middleware.LogHTTP(CreateBoardHandler))

	router.GET("/boards/:bid/tickets", middleware.LogHTTP(GetTicketsHandler))
	router.GET("/boards/:bid/tickets/:tid", middleware.LogHTTP(GetTicketHandler))
	router.POST("/boards/:bid/tickets", middleware.LogHTTP(CreateTicketHandler))
	router.GET("/boards/:bid/newticket", middleware.LogHTTP(GetNewTicketModal))

	router.POST("/users", middleware.LogHTTP(CreateUserHandler))
	router.GET("/users/:uid", middleware.LogHTTP(GetUserHandler))

	if DEBUG {
		router.DELETE("/boards/:bid", RemoveAllBoardsHandler)
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
