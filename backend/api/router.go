package api

import (
	"net/http"
	"os"
	"time"

	"github.com/gonzabosio/jpbabulias/api/handlers"
	"github.com/gonzabosio/jpbabulias/api/mw"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httprate"
)

func NewRouter() (*chi.Mux, error) {
	h, err := handlers.NewHandler()
	if err != nil {
		return nil, err
	}
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{os.Getenv("FRONTEND_URL")},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS", "HEAD"},
		AllowCredentials: true,
	}))
	r.Use(middleware.Logger)
	r.Use(httprate.LimitByIP(100, time.Minute))
	r.Head("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	r.Route("/user", func(r chi.Router) {
		r.Post("/signup", h.UserSignUpHandler)
		r.Post("/login", h.UserLoginHandler)
		r.Post("/logout", h.LogoutHandler)
	})

	r.Group(func(r chi.Router) {
		r.Use(mw.Authenticator)
		r.Route("/appointment", func(r chi.Router) {
			r.Post("/", h.AddAppointmentHandler)
			r.Get("/{user_id}", h.GetAppointmentsByUserIdHandler)
			r.Get("/day", h.GetAppointmentsByDayHandler)
			r.Get("/full", h.GetFullyBookedDatesHandler)
			r.Delete("/{appt_id}", h.DeleteAppointmentHandler)
		})

		r.Route("/patient", func(r chi.Router) {
			r.Get("/{user_id}", h.GetPatientsByUserIdHandler)
			r.Post("/", h.AddPatientHandler)
			r.Patch("/", h.EditPatientDataHandler)
			r.Delete("/{patient_id}", h.DeletePatientByIdHandler)
		})
	})

	r.Route("/bot", func(r chi.Router) {
		r.Post("/prompt", h.SendPromptHandler)
	})
	return r, nil
}
