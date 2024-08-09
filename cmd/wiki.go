package main

import (
	"log"
	configs "webwiki/configa"
	mhttp "webwiki/pkg"
)

func main() {
	server := mhttp.NewServer(configs.SERVER_STATIC_DIRECTORY, configs.SERVER_URL, configs.SERVER_PORT)
	if err := server.InitializeHandlerFunctions(); err != nil {
		log.Fatal(err)
	}

	server.ListenAndServe()
}
