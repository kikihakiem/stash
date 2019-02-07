package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/mux"
	"github.com/kikihakiem/stash/go/simple-crud/common"
	"github.com/kikihakiem/stash/go/simple-crud/controller"
)

func main() {

	db, err := initDB()
	if err != nil {
		log.Fatalf("cannot initialize DB connection: %v", err)
	}

	r := mux.NewRouter()
	// register handlers for /products
	productController := controller.NewProductController(db, log.Printf)
	productController.Handle(r.PathPrefix("/products").Subrouter())

	log.Fatal(http.ListenAndServe(":3000", r))
}

func initDB() (*sql.DB, error) {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	return common.GetDB(dbHost, dbPort, dbUser, dbPass, dbName)
}
