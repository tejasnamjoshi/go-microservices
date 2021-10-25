package handlers

import (
	"go-microservices/product-images/files"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
)

type Files struct {
	store files.Storage
	log   *log.Logger
}

func NewFiles(s files.Storage, l *log.Logger) *Files {
	return &Files{store: s, log: l}
}

func (f Files) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fn := vars["filename"]

	f.log.Println("POST", id, fn)

	f.saveFile(id, fn, rw, r)
}

func (f Files) saveFile(id string, fn string, rw http.ResponseWriter, r *http.Request) {
	f.log.Println("saveFile", id, fn)

	fp := filepath.Join(id, fn)
	err := f.store.Save(fp, r.Body)
	if err != nil {
		f.log.Fatal(err)
		http.Error(rw, "Unable to save the file", http.StatusInternalServerError)
	}
}
