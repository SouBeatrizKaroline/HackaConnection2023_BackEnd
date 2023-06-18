package infra

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/julioc98/citi/internal/app"
	"github.com/julioc98/citi/internal/domain"
)

type ShippingHandler struct {
	usecase *app.ShippingUseCase
	fileDir string
}

func NewShippingHandler(usecase *app.ShippingUseCase, fileDir string) *ShippingHandler {
	return &ShippingHandler{
		usecase: usecase,
		fileDir: fileDir,
	}
}

func NewRouter(r *chi.Mux, h *ShippingHandler) *chi.Mux {
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Post("/upload", h.UploadHandler())
	r.Get("/history", h.HistoryHandler())

	return r
}

func (h *ShippingHandler) UploadHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		file, header, err := r.FormFile("file")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file.Close()

		fileBytes, err := io.ReadAll(file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		filename := header.Filename

		destination := h.fileDir + filename // Set the path where the file will be saved

		newFile, err := os.Create(destination)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer newFile.Close()

		_, err = newFile.Write(fileBytes)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprintln(w, "File saved successfully.")
	}
}

func (h *ShippingHandler) HistoryHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		date := time.Now()

		history := domain.History{
			Shippings: []domain.ShippingHistory{
				{ID: "1", FileName: "REMESSA_CNAB_750.txt", Status: domain.StatusSuccess, Returnpath: "https://9deb-2804-14c-88-22bf-9c6c-3a60-17cc-f9f9.ngrok-free.app/return/RETURN_REMESSA_CNAB_750.txt", CreatedAt: &date, UpdatedAt: nil},
				{ID: "2", FileName: "CNAB750_20230618.txt", Status: domain.StatusProcessing, Returnpath: "", CreatedAt: &date, UpdatedAt: nil},
				{ID: "3", FileName: "CNAB750_20230616.txt", Status: domain.StatusError, Returnpath: "", CreatedAt: &date, UpdatedAt: nil},
			},
		}

		body, err := json.Marshal(history)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(body)
	}
}
