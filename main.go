package main

import (
	api "foodapi/handlers"
	"log"
	"net/http"
)

func main() {
    router := api.SetupRouter()

    log.Fatal(http.ListenAndServe("0.0.0.0:8080", router))
}
