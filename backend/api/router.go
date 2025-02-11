package api

import (
	"net/http"
	"os"
	"time"

	"github.com/gonzabosio/jpbabulias/api/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httprate"
)

func newRouter() (*chi.Mux, error) {
	h, err := handlers.NewHandler()
	if err != nil {
		return nil, err
	}
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{os.Getenv("FRONTEND_URL")},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "HEAD"},
		AllowCredentials: true,
	}))
	r.Use(middleware.Logger)
	r.Use(httprate.LimitByIP(100, time.Minute))

	r.Head("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	r.Route("/appointment", func(r chi.Router) {
		r.Post("/", h.AddAppointmentHandler)
		r.Get("/{user_id}", h.GetAppointmentsByUserIdHandler)
		r.Delete("/{appt_id}", h.DeleteAppointmentHandler)
	})
	return r, nil
}
