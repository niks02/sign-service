package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/signature_service/config"
	"github.com/signature_service/signature"
	"github.com/signature_service/transaction"
)

func main() {
	router := mux.NewRouter().StrictSlash(false)
	port := config.GetHTTPConfig().GetPort()
	config.InitService(router)
	transaction.InitService(router)
	signature.InitService(router)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), router))
}
