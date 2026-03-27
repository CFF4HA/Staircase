package middleware

import (
	"bytes"
	"encoding/json"
	"net/http"
	"slices"

	"github.com/CFF4HA/Staircase/internal/types"
	"github.com/CFF4HA/Staircase/internal/web/core"
	"github.com/DAlba-sudo/pbf"
	"github.com/google/uuid"
)

var (
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

	// creates the session payload we will
	// send to the backend server.
	var session types.Session
	session_id, err := uuid.Parse(c.Value)
	if err != nil {
		// cookie payload is not a valid uuid
		// so we send them back to the
		// login page
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		core.Logger.Info("no session cookie found, redirecting to login page")
		return pbf.ErrSkipHandler
	}

	session.Session = session_id

	// here we prep the request to send to the
	// backend server.
	data, err := json.Marshal(&session)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		core.Logger.Error("failed to marshal session payload, redirecting to login page", "error", err)
		return pbf.ErrSkipHandler
	}

	// here we are actually making the request
	req, err := http.NewRequest(http.MethodPost, core.BackendURL+"/user/session", bytes.NewBuffer(data))
	if err != nil {
		core.Logger.Error("failed to create request to backend server for session validation", "error", err)
		return pbf.ErrSkipHandler
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		core.Logger.Error("failed to create request to backend server for session validation", "error", err)
		return pbf.ErrSkipHandler
	}

	if resp.StatusCode == http.StatusUnauthorized {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		core.Logger.Error("failed to marshal session payload, redirecting to login page", "error", err)
		return pbf.ErrSkipHandler
	} else if resp.StatusCode >= 400 {
		core.Logger.Error("failed to create request to backend server for session validation", "error", err)
		return pbf.ErrSkipHandler
	}

	return nil
}
