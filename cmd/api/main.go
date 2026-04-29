package main

import (
	delivery "backend/internal/delivery/http"
	"backend/internal/repository"
	"backend/internal/usecase"
	"log"
	"net/http"
	"os"
)

func main() {

	repo := &repository.MemoryRepo{}
	summaryUC := &usecase.SummaryUsecase{}

	http.HandleFunc("/init", delivery.CORS(delivery.InitHandler(repo)))
	http.HandleFunc("/summary", delivery.CORS(delivery.SummaryHandler(summaryUC)))
	http.HandleFunc("/orders", delivery.CORS(delivery.OrdersHandler(repo)))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("running :" + port)
	http.ListenAndServe(":"+port, nil)
}