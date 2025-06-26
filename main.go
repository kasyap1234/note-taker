package main

import (
	"context"
	"log"
	"note-taker/api"
	"note-taker/db"
	"note-taker/internal"
	"note-taker/repository"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	// Load .env file. It's good practice to not fail hard if it's missing,
	// as environment variables can be set directly in production.
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found")
	}

	conn, err := db.InitDB()
	if err != nil {
		log.Fatalf("unable to connect to db: %v", err)
	}
	defer conn.Close(context.Background())

	// 1. Initialize Repository (Data Layer)
	queries := repository.New(conn)
	// 2. Initialize Service (Business Logic Layer)
	noteService := internal.NewNoteService(queries)
	// 3. Initialize Handler (Presentation Layer)
	noteHandler := api.NewNoteHandler(noteService)

	e := echo.New()
	e.POST("/notes", noteHandler.CreateNote)
	e.Logger.Fatal(e.Start(":8080"))
}
