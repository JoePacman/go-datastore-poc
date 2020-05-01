package authentication

import (
	"go-datastore-poc/_common"
	"errors"
	"log"
	"net/http"
	"strings"
)

func MustAuth(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if err := checkToken(r); err != nil {
			_common.WriteEmpty(w, http.StatusUnauthorized)
			log.Print(err)
			return
		}

		h.ServeHTTP(w, r) // call original
	})
}

func checkToken(r *http.Request) error {

	_, err := getTokenFromHeader(r)
	if err != nil {
		return err
	}

	// ToDo: Decode JWT token and validate claims
	return nil

	//ToDo: Alternatively: https://cloud.google.com/endpoints/docs/openapi/get-started-app-engine-standard
}

func getTokenFromHeader(r *http.Request) (*string, error) {
	reqToken := r.Header.Get("Authorization")
	if len(reqToken) == 0 {
		return nil, errors.New("authorization token is empty")
	}
	splitToken := strings.Split(reqToken, "Bearer ")
	if len(splitToken) != 2 {
		return nil, errors.New("error splitting authorization token")
	}
	reqToken = splitToken[1]
	if len(reqToken) == 0 {
		return nil, errors.New("authorization token is empty")
	}

	return &reqToken, nil
}
