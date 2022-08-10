package appMiddleware

import (
	"log"
	"net/http"
)

// type appAhandler struct {
// }

func Middleware(next http.HandlerFunc) http.HandlerFunc {
	/*

	 */
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Middleware called")
		r.FormValue("password")
		next(w, r)
	}
}
