package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fmt.Fprint(writer, "Hello GET!")
	})

	server := http.Server{
		Handler: router,
		Addr:    "localhost:3000",
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
