package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"qianbao.com/examples/http-request/handlers"
)

func registerhandler() *httprouter.Router {
	router := httprouter.New()

	router.POST("/user", handlers.TestHandler)

	return router
}

func main() {
	r := registerhandler()
	if err := http.ListenAndServe(":9000", r); err != nil {
		log.Fatalf(err.Error())
	}
}


