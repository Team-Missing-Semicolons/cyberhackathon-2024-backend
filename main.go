package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	_ "github.com/go-sql-driver/mysql"

	"github.com/Team-Missing-Semicolons/cyberhackathon-2024-backend/config"
	"github.com/Team-Missing-Semicolons/cyberhackathon-2024-backend/internal/controller"
	"github.com/Team-Missing-Semicolons/cyberhackathon-2024-backend/internal/mysql"
	"github.com/Team-Missing-Semicolons/cyberhackathon-2024-backend/internal/usecase"
)

// handlePing is a simple ping/pong endpoint handler.
// Used to check if server is alive.
func handlePing(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

// connectToDatabase creates a connection to the database via the given dataSourceName.
func connectToDatabase(dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	// Read in the app configuration.
	cfg, err := config.New(true)
	if err != nil {
		log.Fatal("Failed to load config", err)
	}

	// Connect to the database.
	db, err := connectToDatabase(cfg.Database.DSN)
	if err != nil {
		log.Fatal("failed to connect with the database", err)
	}

	// Setup domain logic.
	logDataStore := mysql.NewLogDataStore(db)
	logUseCase := usecase.NewLogUseCase(logDataStore)
	logController := controller.NewLogController(logUseCase)

	r := chi.NewRouter()

	corsOpts := cors.Handler(cors.Options{
		AllowedOrigins:   cfg.Server.CorsAllowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
	})

	r.Use(corsOpts)

	// Initialize the server multiplexer (router).

	r.Get("/ping", handlePing)
	r.Get("/log", logController.HandleGetLogs)
	r.Post("/log", logController.HandleInsertLog)

	// Create a new http server.
	server := &http.Server{
		Addr:    ":3000",
		Handler: r,
	}

	// Start the server.
	fmt.Println("Starting server at http://localhost:3000")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Server error")
	}
}
