package router

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"io"
	"net/http"

	"github.com/CFF4HA/Staircase/internal/api/core"
	"github.com/CFF4HA/Staircase/internal/api/database"
	"github.com/CFF4HA/Staircase/internal/types"
)

func HandleJobPUT(w http.ResponseWriter, r *http.Request) error {
	db := database.Database()

	var request types.DatabaseJob
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return err
	}

	owner_id, err := getUserIdFromSession(r)
	if err != nil {
		core.Logger.Error("error fetching user id from session", "error", err)
		return err
	}

	uid := uuid.New()
	request.ID = uid
	request.OwnerId = owner_id

	job_tx := db.Create(&request)
	if job_tx.Error != nil {
		core.Logger.Error("error creating job", "error", job_tx.Error)
		return job_tx.Error
	}

	job_meta := &types.DatabaseJobMetadata{
		ID:    uuid.New(),
		JobId: request.ID,
	}
	if tx := db.Create(&job_meta); tx.Error != nil {
		job_tx.Rollback()

		core.Logger.Error("error creating job metadata", "error", tx.Error)
		return tx.Error
	}

	return nil
}

func HandleJobGET(w http.ResponseWriter, r *http.Request) error {
	db := database.Database()
	var request types.DatabaseJob
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		if err != io.EOF {
			core.Logger.Error("error decoding request body", "error", err)
			return err
		}
	}

	owner_id, err := getUserIdFromSession(r)
	if err != nil {
		core.Logger.Error("error fetching user id from session", "error", err)
		return err
	}

	var response any

	if request.ID != uuid.Nil {
		job := &types.DatabaseJob{}
		if tx := db.Where("owner_id = ?", owner_id).Preload("Metadata").Preload("Staircases").First(&job, "id = ?", request.ID); tx.Error != nil {
			core.Logger.Error("error fetching job", "error", tx.Error)
			return tx.Error
		}

		response = job
	} else {
		var jobs []types.DatabaseJob
		if tx := db.Where("owner_id = ?", owner_id).Where(&request).Preload("Metadata").Preload("Staircases").Find(&jobs); tx.Error != nil {
			core.Logger.Error("error fetching jobs", "error", tx.Error)
			return tx.Error
		}

		response = jobs
	}

	return json.NewEncoder(w).Encode(response)
}

func HandleJobDELETE(w http.ResponseWriter, r *http.Request) error {
	db := database.Database()
	var request types.DatabaseJob
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		core.Logger.Error("error decoding request body", "error", err)
		return err
	} else if request.ID == uuid.Nil {
		core.Logger.Error("id is required for delete")
		return fmt.Errorf("id is required for delete")
	}

	owner, err := getUserIdFromSession(r)
	if err != nil {
		core.Logger.Error("error fetching user id from session", "error", err)
		return err
	}

	tx := db.Delete(&types.DatabaseJob{}, "id = ?", request.ID, "owner_id = ?", owner)
	if tx.Error != nil {
		core.Logger.Error("error deleting job", "error", tx.Error)
		return tx.Error
	}

	tx = db.Delete(&types.DatabaseStaircase{}, "job_id = ?", request.ID, "owner_id = ?", owner)
	if tx.Error != nil {
		core.Logger.Error("error deleting staircases", "error", tx.Error)
		return tx.Error
	}

	w.WriteHeader(http.StatusOK)
	return nil
}
