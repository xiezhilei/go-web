
package main

import (
	"net/http"
	"fmt"
)

type HelloHandler struct {

}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "World!")
}

type WorldHandler struct {

}

func (h *WorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "World!")
}

func nihao(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "nihao!")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	hello := HelloHandler{}
	world := WorldHandler{}

	http.Handle("/hello", &hello)
	http.Handle("/world", &world)
	http.HandleFunc("/nihao", nihao)

	server.ListenAndServeTLS("cert.pem", "key.pem")
	//server.ListenAndServe()
}
