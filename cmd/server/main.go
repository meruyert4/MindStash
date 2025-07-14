package main

import (
	"context"
	"database/sql"
	"log"
	"mindstash/internal/note"
	"mindstash/pkg/logger"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	ctx := context.Background()
	logg := logger.InitLogger(ctx, logger.LevelDebug)

	dsn := "postgres://postgres:postgres@db:5432/mindstash?sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		logg.Error(ctx, "Failed to connect to DB", "error", err)
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		logg.Error(ctx, "DB unreachable", "error", err)
		log.Fatal(err)
	}
	logg.Info(ctx, "Successfully connected to DB")

	repo := note.NewRepository(db, logg)
	service := note.NewService(repo, logg)
	handler := note.NewHandler(service, logg)

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("./frontend")))
	mux.HandleFunc("POST /notes", handler.CreateNote)
	mux.HandleFunc("GET /notes", handler.GetNotes)
	mux.HandleFunc("GET /notes/{id}", handler.GetNoteById)
	mux.HandleFunc("PUT /notes/{id}", handler.UpdateNote)
	mux.HandleFunc("DELETE /notes/{id}", handler.DeleteNote)

	logg.Info(ctx, "Server starting on http://localhost:8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		logg.Error(ctx, "Failed to start server", "error", err)
		log.Fatal(err)
	}
}
