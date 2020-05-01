package datastore_film

import (
	"cloud.google.com/go/datastore"
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"go-datastore-poc/_common"
	"go-datastore-poc/dto"
	"log"
	"net/http"
)

func GetFilm(w http.ResponseWriter, r *http.Request) {

	// Create all the required objects
	ctx, filmService, err := createObjects()
	if checkError(err, w) {
		return
	}
	films, err := filmService.FindByTitle(ctx, mux.Vars(r)["title"])
	if checkError(err, w) {
		return
	}

	_common.WriteJson(w, http.StatusOK, films)
}

func PostFilm(w http.ResponseWriter, r *http.Request) {

	// Decode json
	var f dto.Film
	err := json.NewDecoder(r.Body).Decode(&f)

	// Create all the required objects
	ctx, filmService, err2 := createObjects()

	if checkError(err, w) && checkError(err2, w) {
		return
	}

	if checkError(filmService.Create(ctx, &f), w) {
		return
	}

	_common.WriteEmpty(w, http.StatusOK)
}

func checkError(err error, w http.ResponseWriter) bool {
	if err != nil {
		log.Print(err)
		_common.WriteEmpty(w, http.StatusInternalServerError)
		return true
	} else {
		return false
	}
}

func createObjects() (context.Context, *FilmService, error) {
	// TODO: introduce .properties file to store project ID per environment
	ctx := context.Background()
	datastoreClient, err := datastore.NewClient(ctx, "joe-gcp-playground")
	repo := NewDatastoreRepo(datastoreClient)
	filmService := NewService(repo)
	return ctx, filmService, err
}
