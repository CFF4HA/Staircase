package bridges

import (
	"encoding/json"
	"github.com/CFF4HA/Staircase/internal/types"
	"github.com/google/uuid"
	"net/http"
)

type BridgeJob struct{}

func (b BridgeJob) Data(w http.ResponseWriter, r *http.Request) (any, error) {
	var request types.DatabaseJob

	id, err := uuid.Parse(r.FormValue("id"))
	if err != nil {
		return nil, err
	}
	request.ID = id

	data, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	session, err := r.Cookie("session")
	if err != nil {
		return nil, err
	}

	resp, err := sendToBackend(
		"/job", data, http.MethodGet,
		session,
	)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var job types.DatabaseJob
	if err := json.NewDecoder(resp.Body).Decode(&job); err != nil {
		return nil, err
	}

	return job, nil
}

type BridgeJobs struct{}

func (b BridgeJobs) Data(w http.ResponseWriter, r *http.Request) (any, error) {
	session, err := r.Cookie("session")
	if err != nil {
		return nil, err
	}

	resp, err := sendToBackend(
		"/job", nil, http.MethodGet,
		session,
	)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var job []types.DatabaseJob
	if err := json.NewDecoder(resp.Body).Decode(&job); err != nil {
		return nil, err
	}

	return job, nil

}
