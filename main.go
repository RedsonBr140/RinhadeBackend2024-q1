package main

import (
	"context"
	"log"
	"net/http"

	"os"

	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
)

func main() {
	// urlExample := "postgres://user:pass@localhost:5432/db_name"
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Connecting to the database: %v\n", err)
	}

	defer conn.Close(context.Background())

	// Create router
	e := echo.New()
	// Register routes
	e.GET("/", HelloHandler)
	// Start router
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}

func HelloHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
