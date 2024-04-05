package main

import (
	"context"
	"database/sql"
	"fmt"
	"hello-world/internal/data"
	"net/http"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type application struct {
	models data.Models
}

func main() {
	db, err := openDB()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer db.Close()

	fmt.Println("database connection successful")

	app := &application{
		models: data.NewModels(db),
	}

	server := &http.Server{
		Addr:    ":3000",
		Handler: app.routes(),
	}

	fmt.Println("server running on port :3000")
	server.ListenAndServe()

}

func openDB() (*sql.DB, error) {
	// Here, give your own Postgres credentials
	db, err := sql.Open("pgx", "postgres://akbar:Password123!@localhost:5432/moviesdb")
	if err != nil {
		return nil, err
	}

	newContext := context.Background()

	ctx, cancel := context.WithTimeout(newContext, 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
