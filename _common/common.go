// Commonly used functions
package _common

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func WriteEmpty(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
	if _, err := fmt.Fprint(w); err != nil {
		writeError(w, err)
		return
	}
}

func WriteJson(w http.ResponseWriter, status int, v interface{}) {
	js, err := json.Marshal(v)
	if err != nil {
		writeError(w, err)
		return
	}

	w.WriteHeader(status)
	if _, err := w.Write(js); err != nil {
		writeError(w, err)
		return
	}
}

func writeError(w http.ResponseWriter, err error) {
	log.Print(err)
	http.Error(w, err.Error(), http.StatusInternalServerError)
}
