package router

import (
	"errors"
	"net/http"

	"github.com/CFF4HA/Staircase/internal/api/core"
	"github.com/google/uuid"
)

var (
	backend_url string
)

func SetBackendUrl(url string) {
	backend_url = url
}

func getUserIdFromSession(r *http.Request) (uuid.UUID, error) {
	var nil_uuid uuid.UUID

	c, err := r.Cookie("session")
	if err != nil {
		core.Logger.Error("error fetching session cookie", "error", err)
		return nil_uuid, err
	}

	session, err := uuid.Parse(c.Value)
	if err != nil {
		core.Logger.Error("error parsing session cookie", "error", err)
		return nil_uuid, err
	}

	user, ok := Sessions[session]
	if !ok {
		core.Logger.Error("session not found", "session", session)
		return nil_uuid, errors.New("session not found")
	}

	return user, nil
}
