// SetMiddleWareJSON - this will format all responses to JSON
// SetMiddleWareAuthentication - this will check for the validity of the authentication token provided

package middlewares

import (
	"errors"
	"net/http"

	"github.com/REST_API/api/auth"
	"github.com/REST_API/api/responses"
)

func SetMiddlewareJSON(next http.HandleFunc) http.HandleFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}

func SetMiddlewareAuthentication(next http.HandleFunc) http.HandleFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := auth.TokenValid(r)
		if err != nil {
			responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}
		next(w, r)
	}

}
