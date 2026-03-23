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

type JobPutRequest struct {
	DatabaseConnString string                    `json:"database_conn_string"`
	DatabaseTableName  string                    `json:"database_table_name"`
	DatabaseColumnName string                    `json:"database_column_name"`
	Staircases         []types.DatabaseStaircase `json:"staircases"`
}

func HandleJobPUT(w http.ResponseWriter, r *http.Request) error {
	db := database.Database()

	var request JobPutRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return err
	}

	var job types.DatabaseJob
	job.ID = uuid.New()
	for _, staircase := range request.Staircases {
		job.Staircases = append(job.Staircases, types.DatabaseStaircase{
			ID:          uuid.New(),
			JobId:       job.ID,
			Declaration: staircase.Declaration,
		})
	}

	return db.Create(&job).Error
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

	var response any

	if request.ID != uuid.Nil {
		job := &types.DatabaseJob{}
		if tx := db.Preload("Staircases").First(&job, "id = ?", request.ID); tx.Error != nil {
			core.Logger.Error("error fetching job", "error", tx.Error)
			return tx.Error
		}

		response = job
	} else {
		var jobs []types.DatabaseJob
		if tx := db.Where(&request).Preload("Staircases").Find(&jobs); tx.Error != nil {
			core.Logger.Error("error fetching jobs", "error", tx.Error)
			return tx.Error
		}

		response = jobs
	}

	return json.NewEncoder(w).Encode(response)
}

func HandleJobPATCH(w http.ResponseWriter, r *http.Request) error {
	db := database.Database()
	var request types.DatabaseJob
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		core.Logger.Error("error decoding request body", "error", err)
		return err
	} else if request.ID == uuid.Nil {
		core.Logger.Error("id is required for delete")
		return fmt.Errorf("id is required for delete")
	}

	db.Model(&types.DatabaseJob{}).
		Where("id = ?", request.ID).
		Select(
			"IsInitialized", "IsInProgress", "IsFinished", "IsUploaded",
			"DatabaseConnString", "DatabaseTableName", "DatabaseColumnName",
		).
		Updates(&request)

	tx := db.Model(&types.DatabaseJob{}).Where("id = ?", request.ID).Updates(request)
	if tx.Error != nil {
		core.Logger.Error("error updating job", "error", tx.Error)
		return tx.Error
	}

	return nil
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

	tx := db.Delete(&types.DatabaseJob{}, "id = ?", request.ID)
	if tx.Error != nil {
		core.Logger.Error("error deleting job", "error", tx.Error)
		return tx.Error
	}

	tx = db.Delete(&types.DatabaseStaircase{}, "job_id = ?", request.ID)
	if tx.Error != nil {
		core.Logger.Error("error deleting staircases", "error", tx.Error)
		return tx.Error
	}

	w.WriteHeader(http.StatusOK)
	return nil
}
