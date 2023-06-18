package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/go-chi/chi"
	"github.com/julioc98/citi/internal/app"
	"github.com/julioc98/citi/internal/infra"
)

func main() {
	chiRouter := chi.NewRouter()

	fileDir, err := os.Getwd()
	if err != nil {
		log.Println("Error getting current working directory:", err)
		return
	}

	fileDir = filepath.Join(fileDir, "files")

	// Repository and Gateway
	shippingRepo := infra.NewShippingRepository()
	bacenGateway := infra.NewBacenGateway()

	// Storage
	returnStorage := infra.NewReturnStorage(fileDir)

	// UseCase
	shippingUC := app.NewShippingUseCase(shippingRepo, bacenGateway, returnStorage)

	// Worker
	shippingWorker := infra.NewShippingWorker(shippingUC, fileDir)

	// Handler
	shippingHandler := infra.NewShippingHandler(shippingUC, fileDir)

	r := infra.NewRouter(chiRouter, shippingHandler)

	// Run server
	go runApiHTTP(r) // Run server in a goroutine

	// Run worker
	go runWorker(shippingWorker) // Run worker in an another goroutine

	// Keep the main goroutine running
	select {}

}

func runApiHTTP(handler http.Handler) {

	port := os.Getenv("PORT")
	log.Println("Server running on port", port)
	http.ListenAndServe(":"+port, handler)
}

func runWorker(shippingWorker *infra.ShippingWorker) {
	log.Println("Worker running...")
	for {
		shippingWorker.Run()

		time.Sleep(5 * time.Second) // Check for new files every 5 seconds
	}
}
