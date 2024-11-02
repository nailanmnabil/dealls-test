package main

import (
	"dealls-test/config"
	"dealls-test/handlers"
	"dealls-test/pkg"
	"dealls-test/services"
	"log/slog"
	"net/http"

	_ "dealls-test/docs"

	"github.com/go-chi/chi/v5"
)

// @title Dealls API Doc
// @version 1.0
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description This is a documentation of Dealls backend service.
// @termsOfService http://swagger.io/terms/
// @BasePath /
func main() {
	cfg := config.LoadConfig(".env")
	dbConn := pkg.GetDbConn(cfg.Database.StringConn, cfg.Database.MinConn, cfg.Database.MaxConn)
	svc := services.NewService(dbConn, cfg)

	r := chi.NewRouter()
	handlers.RegisterHandler(r, svc, cfg)

	slog.Info("Server running on port 8080")
	http.ListenAndServe(":8080", r)
}
