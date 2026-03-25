package bridges

import (
	"encoding/json"
	"net/http"

	"github.com/CFF4HA/Staircase/internal/types"
	"github.com/CFF4HA/Staircase/internal/web/core"
)

type BridgeCardList struct{}

func (b BridgeCardList) Data(w http.ResponseWriter, r *http.Request) (any, error) {
	session, err := r.Cookie("session")
	if err != nil {
		core.Logger.Error("no session cookie available", "err", err)
		return nil, err
	}

	// we send a request to the backend server so we can
	// collect all the relevant datasources for the user.
	resp, err := sendToBackend("/ds", nil, http.MethodGet, session)
	if err != nil {
		core.Logger.Error("failed to send request to backend server", "err", err, "session", session)
		return nil, err
	}

	// we now deserialize the response body into
	// a list of data sources
	var sources []types.DatabaseDatasource
	if err := json.NewDecoder(resp.Body).Decode(&sources); err != nil {
		core.Logger.Error("failed to decode response body from backend server", "err", err)
		return nil, err
	}

	return sources, nil
}
