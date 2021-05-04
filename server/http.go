package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/TechharaInc/Yfitops/client"
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
	sc := client.NewSpotifyClient()
	q := r.URL.Query()
	var code string
	for k, v := range q {
		if k == "code" {
			code = v[0]
		}
	}
	if code == "" {
		http.Error(w, "OTOTOI KIYAGARE!", http.StatusNotFound)
		return
	}

	sc.Exchange(code)
	fmt.Fprintf(w, "YOU CAN LEAVE THIS PAGE SAFELY.")
}
