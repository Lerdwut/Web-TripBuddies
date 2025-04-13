package main

import (
    "context"
    "fmt"
    "github.com/jackc/pgx/v5/pgxpool"
    "log"
    "os"
    "github.com/gofiber/fiber/v2"
    "Backend/routes"
)

var DB *pgxpool.Pool

func main() {
    dbURL := os.Getenv("SUPABASE_DB_URL")

    var err error
    DB, err = pgxpool.New(context.Background(), dbURL)
    if err != nil {
        log.Fatal("Can't Connect tp Supabase DB:", err)
    }

    fmt.Println("Connect to Supabase Success")

    app := fiber.New()
    routes.SetupRoutes(app, DB)
    app.Listen(":3000")
}
