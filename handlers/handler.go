package handlers

import (
	"dealls-test/config"
	"dealls-test/services"
	"fmt"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth/v5"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Handler struct {
	svc *services.Service
	cfg *config.Config
}

func RegisterHandler(m *chi.Mux, svc *services.Service, cfg *config.Config) {
	h := &Handler{svc, cfg}

	m.Use(middleware.Logger)
	m.Use(middleware.Recoverer)
	m.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(fmt.Sprintf("http://%s:%s/swagger/docs/doc.json", cfg.App.Host, cfg.App.Port)),
	))

	// public
	m.Post("/register", h.Register)
	m.Post("/login", h.Login)
	m.Get("/packages", h.GetAllPackage)

	// protected
	m.Group(func(r chi.Router) {
		tokenAuth := jwtauth.New("HS256", []byte(cfg.Jwt.AccessTokenSecret), nil)
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator(tokenAuth))

		r.Get("/profiles/random", h.GetRandomProfile)
		r.Post("/profiles/swipe", h.Swipe)
		r.Post("/packages/purchase", h.PurchasePackage)
	})
}
