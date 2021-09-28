package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		http.Error(rw, "Oops", 500)
		// rw.Write([]byte("Oops"));
		return
	}
	h.l.Println("Hello", d)
	fmt.Fprintf(rw, "Hello Resp %s", d)
}
