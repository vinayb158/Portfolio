package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

func main() {
	DbConn := os.Getenv("Portfolio_DB_Connection")
	if DbConn == "" {
		fmt.Println("DB connection string is missing!")
		return
	}

	fmt.Println("Using DB Connection:", DbConn)
	conn, err := pgx.Connect(context.Background(), DbConn)
    if err != nil {
        log.Fatalf("Unable to connect to database: %v\n", err)
    }
    defer conn.Close(context.Background())

    // Run a simple query
    rows, err := conn.Query(context.Background(), "SELECT party,secu FROM porttrxn")
    if err != nil {
        log.Fatalf("Query failed: %v\n", err)
    }
    defer rows.Close()

    // Process the result
    for rows.Next() {
        var party string
        var secu string
        if err := rows.Scan(&party, &secu); err != nil {
            log.Fatalf("Row scan failed: %v\n", err)
        }

        fmt.Printf("Party: %s, Scrip: %s\n", party, secu)
    }
    // Check if there were any errors in the row iteration
    if err := rows.Err(); err != nil {
        log.Fatalf("Rows iteration failed: %v\n", err)
    }
}
