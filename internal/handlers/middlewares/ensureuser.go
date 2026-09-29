package middlewares

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"gopkg.in/telebot.v4"
)

const upsertUserQuery = `
INSERT INTO users (id, first_name, last_name, username, last_activity_at)
VALUES ($1, $2, NULLIF($3, ''), NULLIF($4, ''), NOW())
ON CONFLICT (id) DO UPDATE SET
	first_name       = EXCLUDED.first_name,
	last_name        = EXCLUDED.last_name,
	username         = EXCLUDED.username,
	last_activity_at = NOW()
`

func EnsureUser(pool *pgxpool.Pool) telebot.MiddlewareFunc {
	return func(next telebot.HandlerFunc) telebot.HandlerFunc {
		return func(c telebot.Context) error {
			sender := c.Sender()
			if sender == nil || sender.IsBot {
				return next(c)
			}

			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			_, err := pool.Exec(ctx, upsertUserQuery,
				sender.ID,
				sender.FirstName,
				sender.LastName,
				sender.Username,
			)
			if err != nil {
				return fmt.Errorf("ensure user %d: %w", sender.ID, err)
			}

			return next(c)
		}
	}
}