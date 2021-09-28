package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Goodbye struct {
	l *log.Logger
}

func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

func (g *Goodbye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Oops from goodbye", http.StatusInternalServerError)
		return
	}
	g.l.Println("Goodbye", d)
	fmt.Fprintf(rw, "Goodbye %s", d)
}
