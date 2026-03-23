package router

import (
	"net/http"

	"github.com/CFF4HA/Staircase/internal/api/core"
	"github.com/DAlba-sudo/pbf"
)

func CookiePrinter(w http.ResponseWriter, r *http.Request) error {
	core.Logger.Info("cookies", "cookies", r.Cookies())

	return nil
}

func CORS(w http.ResponseWriter, r *http.Request) error {
	if backend_url == "" {
		w.Header().Set("Access-Control-Allow-Origin", "*")
	} else {
		w.Header().Set("Access-Control-Allow-Origin", backend_url)
	}
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	if r.Method != http.MethodOptions {
		return nil
	}

	w.WriteHeader(http.StatusNoContent)
	return pbf.ErrSkipHandler
}
