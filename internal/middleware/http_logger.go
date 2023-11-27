package middleware

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
)

type statusInterceptor struct {
	http.ResponseWriter
	status int
}

func (intr *statusInterceptor) WriteHeader(code int) {
	intr.status = code
	intr.ResponseWriter.WriteHeader(code)
}

func LogHTTP(n httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		endpoint := fmt.Sprintf("%s %s", strings.ToUpper(r.Method), r.URL)
		log.Printf(endpoint)
		i := &statusInterceptor{w, 200}
		n(i, r, p)

		log.Printf("%d - %s", i.status, endpoint)
	}
}
