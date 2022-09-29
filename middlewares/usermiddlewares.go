package middlewares

import (
	"github.com/rs/cors"
	"github.com/uptrace/bunrouter"
)

func NewCorsMiddleware(next bunrouter.HandlerFunc) bunrouter.HandlerFunc {
	corsHandler := cors.New(cors.Options{
		AllowCredentials: true,
		AllowedHeaders:   []string{"Access-Control-Allow-Origin", "Content-Type", "Origin", "Accept"},
	})
	return bunrouter.HTTPHandler(corsHandler.Handler(next))
}
