// cmd/server/main.go
package main

import (
	"fmt"
	"log"
	"net/http"
	"MME/internal/repository"
	"MME/internal/service"
)

func main() {
	port := ":8080"
	dbFile := "./metadata.db"

	fmt.Printf("🚀 Microservice booting up. Storage backend: SQLite (%s)\n", dbFile)

	// 1. Initialize Database Repository
	dbRepo, err := repository.NewDbRepository(dbFile)
	if err != nil {
		log.Fatalf("❌ Failed to initialize database: %v", err)
	}
	defer dbRepo.Close()

	// 2. Dependency Injection
	extractor := service.NewExtractorService()
	handler := service.NewHttpHandler(extractor, dbRepo)

	// 3. Routing
	http.HandleFunc("/extract", handler.ExtractHandler)

	log.Fatal(http.ListenAndServe(port, nil))
}