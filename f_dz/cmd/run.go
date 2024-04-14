package cmd

import (
	"f_dz/db"
	"f_dz/handlers"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func Run() {
	fmt.Println("Starting the database...")
	db.InitDB()
	db.Connect()
	fmt.Println("Connected to the database")
	db.Migrate()

	r := mux.NewRouter()

	r.HandleFunc("/commissions/calculate", handlers.CalculateCommission)
	r.HandleFunc("/transactions", handlers.HandleTransactions)
	r.HandleFunc("/transactions", handlers.Authenticate(handlers.HandleTransactions)).Methods("GET", "POST")
	r.HandleFunc("/transactions/{id}", handlers.HandleTransactions)
	r.HandleFunc("/users", handlers.RegisterUser).Methods("POST")
	r.HandleFunc("/users/login", handlers.LoginUser).Methods("POST")

	http.Handle("/", r)

	fmt.Println("Server is running on port 8080")
	fmt.Println("Press Ctrl+C to quit.")
	log.Fatal(http.ListenAndServe(":8080", nil))

	var log = logrus.New()
	// Настройка формата логов
	log.Formatter = &logrus.JSONFormatter{}
	// Установка минимального уровня логирования
	log.Level = logrus.InfoLevel
}
