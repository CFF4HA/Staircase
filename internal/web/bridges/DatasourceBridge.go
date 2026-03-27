package bridges

import (
	"encoding/json"
	"github.com/CFF4HA/Staircase/internal/types"
	"github.com/google/uuid"
	"net/http"
)

type BridgeDatasource struct{}

func (b BridgeDatasource) Data(w http.ResponseWriter, r *http.Request) (any, error) {
	session, err := r.Cookie("session")
	if err != nil {
		return nil, err
	}

	u, err := uuid.Parse(r.FormValue("id"))
	if err != nil {
		return nil, err
	}

	// craft the request with the ID as part of the object
	request := &types.DatabaseDatasource{
		ID: u,
	}
	data, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	// we send a request to the backend server so we can
	// collect all the relevant datasources for the user.
	resp, err := sendToBackend("/ds", data, http.MethodGet, session)
	if err != nil {
		return nil, err
	}

	// we now deserialize the response body into
	// a list of data sources
	var source types.DatabaseDatasource
	if err := json.NewDecoder(resp.Body).Decode(&source); err != nil {
		return nil, err
	}

	return source, nil
}

type BridgeDatasources struct{}

func (b BridgeDatasources) Data(w http.ResponseWriter, r *http.Request) (any, error) {
	session, err := r.Cookie("session")
	if err != nil {
		return nil, err
	}

	// we send a request to the backend server so we can
	// collect all the relevant datasources for the user.
	resp, err := sendToBackend("/ds", nil, http.MethodGet, session)
	if err != nil {
		return nil, err
	}

	// we now deserialize the response body into
	// a list of data sources
	var sources []types.DatabaseDatasource
	if err := json.NewDecoder(resp.Body).Decode(&sources); err != nil {
		return nil, err
	}

	return sources, nil
}
