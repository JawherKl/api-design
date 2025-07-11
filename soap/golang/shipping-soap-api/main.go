package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/JawherKl/shipping-soap-api/api/soap"
	"github.com/JawherKl/shipping-soap-api/storage"
)

func main() {
	store := storage.NewStore()
	orderService := &soap.OrderService{Store: store}

	http.Handle("/soap", orderService)

	fmt.Println("🚀 SOAP API running on http://localhost:8080/soap")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
