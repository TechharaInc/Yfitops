package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/TechharaInc/Yfitops/client"
	"github.com/TechharaInc/Yfitops/service"
)

type HttpServer struct{}

func NewHttpServer() *HttpServer {
	http.HandleFunc("/callback", completeAuth)
	go func() {
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatal(err)
		}
	}()

	log.Println("Server start listening at :8080")

	return &HttpServer{}
}

func completeAuth(w http.ResponseWriter, r *http.Request) {
	sc, err := client.NewSpotifyClient()
	if err != nil {
		http.Error(w, ":weary:", http.StatusInternalServerError)
		return
	}

	q := r.URL.Query()
	var code string
	var state string
	for k, v := range q {
		if k == "code" {
			code = v[0]
		}
		if k == "state" {
			state = v[0]
		}
	}
	if code == "" {
		http.Error(w, ":weary:", http.StatusBadRequest)
		return
	}

	tok, err := sc.Exchange(code)
	if err != nil {
		http.Error(w, ":weary:", http.StatusInternalServerError)
		return
	}

	service.SetGuildIDToContext(r.Context(), state)
	service.SetTokenToContext(r.Context(), tok)

	fmt.Fprintf(w, ":ok::ok:")
}
