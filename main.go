package main

import (
	"log"
	"net/http"

	"inventoryService/product"
)

const basePath = "/api"

func main() {
	product.SetupRoutes(basePath)
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
