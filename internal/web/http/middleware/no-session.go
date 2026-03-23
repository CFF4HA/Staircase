package middleware

import (
	"net/http"
	"slices"

	"github.com/CFF4HA/Staircase/internal/api/core"
	"github.com/DAlba-sudo/pbf"
)

var (
	Sessions = make(map[string]bool)

	ExcludedSessionEndpoints = []string{
		"/login",
		"/htmx/form/login",
		"/htmx/form/signup",
	}
)

func MwNoSession(w http.ResponseWriter, r *http.Request) error {
	if slices.Contains(ExcludedSessionEndpoints, r.URL.Path) {
		core.Logger.Info("endpoint is excluded from session check, skipping middleware", "endpoint", r.URL.Path)
		return nil
	}

	c, err := r.Cookie("session")
	if err != nil {
		if err == http.ErrNoCookie {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			core.Logger.Info("no session cookie found, redirecting to login page")
			return pbf.ErrSkipHandler
		}
	}

	_, ok := Sessions[c.Value]
	if !ok {
		Sessions[c.Value] = true
	}

	return nil
}
