package main

import (
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Print("API Rodando!!")
	r := router.Gerar()
	log.Fatal(http.ListenAndServe(":5001", r))
}
