package datastore_film

import (
	"cloud.google.com/go/datastore"
	"context"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"go-datastore-poc/_common"
	"go-datastore-poc/dto"
	"log"
	"net/http"
)

func GetFilm(w http.ResponseWriter, r *http.Request) {

	// Create all the required objects
	ctx, filmService, err := createObjects()

	films, err := filmService.FindByTitle(ctx, mux.Vars(r)["title"])

	if err != nil {
		log.Print(err)
		_common.WriteEmpty(w, http.StatusInternalServerError)
	} else {
		_common.WriteJson(w, http.StatusOK, films)
	}
}

func PostFilm(w http.ResponseWriter, r *http.Request) {

	// Create all the required objects
	ctx, filmService, err := createObjects()
	err = errors.New("my error")
	checkError(err, w)


	var f dto.Film
	err = json.NewDecoder(r.Body).Decode(&f)

	err = filmService.Create(ctx, &f)


}

func checkError(err error, w http.ResponseWriter) {
	if err != nil {
		log.Print(err)
		_common.WriteEmpty(w, http.StatusInternalServerError)
	} else {
		_common.WriteEmpty(w, 200)
	}
}

func createObjects() (context.Context, *FilmService, error)  {
	// TODO: introduce .properties file to store project ID per environment
	ctx := context.Background()
	datastoreClient, err := datastore.NewClient(ctx, "joe-gcp-playground")
	repo := NewDatastoreRepo(datastoreClient)
	filmService := NewService(repo)
	return ctx, filmService, err
}

