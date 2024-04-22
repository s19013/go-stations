package router

import (
	"database/sql"
	"net/http"

	"github.com/TechBowl-japan/go-stations/controller"
)

func NewRouter(todoDB *sql.DB) *http.ServeMux {
	// register routes
	mux := http.NewServeMux()

	mux.HandleFunc("/healthz", controller.Healthz)

	return mux
}
