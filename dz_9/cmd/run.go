package cmd

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/DedFura/crud-example/handlers"
)

func Run(db *sql.DB) {
	http.HandleFunc("/item", handlers.Item(db))
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}