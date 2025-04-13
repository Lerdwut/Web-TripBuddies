package routes

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

type User struct {
	ID        int       `json:"id"`
	GoogleID  string    `json:"google_id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	AvatarURL string    `json:"avatar_url"`
}

func SetupRoutes(app *fiber.App, DB *pgxpool.Pool) {
    app.Get("/users/v1", func(c *fiber.Ctx) error {
        rows, err := DB.Query(context.Background(), "SELECT id, google_id, email, name, avatar_url FROM users")
        if err != nil {
            return c.Status(500).SendString(fmt.Sprintf("Scan error: %v", err))
        }
        defer rows.Close()
        
        var users []User
        for rows.Next() {
            var user User
            if err := rows.Scan(&user.ID, &user.GoogleID, &user.Email, &user.Name, &user.AvatarURL); err != nil {
                return c.Status(500).SendString(fmt.Sprintf("Scan error: %v", err))
            }
            users = append(users, user)
        }
        return c.JSON(users)
    })    

    
}