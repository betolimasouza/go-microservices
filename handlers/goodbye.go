package handlers

import (
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
	g.l.Println("Goodbye World!")
	d, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(rw, "Error on Goodbye!", http.StatusBadRequest)
		return
	}

	g.l.Printf("Bye %s!", string(d))
}
