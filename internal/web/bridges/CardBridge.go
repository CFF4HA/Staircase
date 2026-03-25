package bridges

import (
	"encoding/json"
	"net/http"

	"github.com/CFF4HA/Staircase/internal/types"
	"github.com/CFF4HA/Staircase/internal/web/core"
	"github.com/google/uuid"
)

type BridgeCard struct{}

func (b BridgeCard) Data(w http.ResponseWriter, r *http.Request) (any, error) {
	session, err := r.Cookie("session")
	if err != nil {
		core.Logger.Error("no session cookie available", "err", err)
		return nil, err
	}

	u, err := uuid.Parse(r.FormValue("id"))
	if err != nil {
		core.Logger.Error("invalid uuid provided", "err", err, "id", r.FormValue("id"))
		return nil, err
	}

	// craft the request with the ID as part of the object
	request := &types.DatabaseDatasource{
		ID: u,
	}
	data, err := json.Marshal(request)
	if err != nil {
		core.Logger.Error("failed to marshal request body", "err", err, "request", request)
		return nil, err
	}

	// we send a request to the backend server so we can
	// collect all the relevant datasources for the user.
	resp, err := sendToBackend("/ds", data, http.MethodGet, session)
	if err != nil {
		core.Logger.Error("failed to send request to backend server", "err", err, "session", session)
		return nil, err
	}

	// we now deserialize the response body into
	// a list of data sources
	var source types.DatabaseDatasource
	if err := json.NewDecoder(resp.Body).Decode(&source); err != nil {
		core.Logger.Error("failed to decode response body from backend server", "err", err)
		return nil, err
	}

	w.Header().Set("Hx-Retarget", "#modal-content-target")
	w.Header().Set("Hx-Trigger", "showModal")
	return source, nil
}
