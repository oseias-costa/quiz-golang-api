package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type muxRouter struct{}

var muxDispacther = mux.NewRouter()

func NewMuxRouter() Router {
	return &muxRouter{}
}

func (*muxRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispacther.HandleFunc(uri, f).Methods(http.MethodGet)
}

func (*muxRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispacther.HandleFunc(uri, f).Methods(http.MethodPost)
}

func (*muxRouter) SERVE(port string) {
	fmt.Printf("Mux HTTP server runnig on port %s", port)
	http.ListenAndServe(port, muxDispacther)
}
