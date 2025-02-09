package main

import (
	"fmt"
	handlers "github.com/DaniilStelmakh/backend/src/api"
	"github.com/DaniilStelmakh/backend/src/database"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	// Загружаем переменные окружения из файла .env
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	err = database.CreateTable(db, ".")
	if err != nil {
		log.Fatal(err)
	}

	// Новый роутер
	r := chi.NewRouter()

	// Добавляем обработчики для маршрутов /pings
	r.Get("/pings", handlers.ShowPingsHandler(db))
	r.Post("/pings", handlers.WritePingHandler(db))

	// Запускаем сервер на указанном порту
	fmt.Printf("Listening on port %s\n", port)
	http.ListenAndServe(":"+port, r)

}
