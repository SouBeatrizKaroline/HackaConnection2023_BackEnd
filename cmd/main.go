package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/julioc98/citi/internal/app"
	"github.com/julioc98/citi/internal/infra"
)

func main() {
	chiRouter := chi.NewRouter()

	// Repository and Gateway
	shippingRepo := infra.NewShippingRepository()
	bacenGateway := infra.NewBacenGateway()

	// Storage
	returnStorage := infra.NewReturnStorage()

	// UseCase
	shippingUC := app.NewShippingUseCase(shippingRepo, bacenGateway, returnStorage)

	fileDir := "/Users/jc/Projects/github.com/julioc98/citi/files/"

	// Worker
	shippingWorker := infra.NewShippingWorker(shippingUC, fileDir)

	// Handler
	shippingHandler := infra.NewShippingHandler(shippingUC, fileDir)

	r := infra.NewRouter(chiRouter, shippingHandler)

	// Run server
	go runApiHTTP(r)

	// Run worker
	go runWorker(shippingWorker)

	// Keep the main goroutine running
	select {}

}

func runApiHTTP(handler http.Handler) {

	port := os.Getenv("PORT")
	// port := "8001"
	log.Println("Server running on port", port)
	// http.ListenAndServe(":"+port, handler)
}

func runWorker(shippingWorker *infra.ShippingWorker) {
	log.Println("Worker running...")
	for {
		shippingWorker.Run()

		time.Sleep(5 * time.Second) // Check for new files every 5 seconds
	}
}
