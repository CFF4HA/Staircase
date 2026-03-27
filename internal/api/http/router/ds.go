package router

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/google/uuid"

	"github.com/CFF4HA/Staircase/internal/api/core"
	"github.com/CFF4HA/Staircase/internal/api/database"
	"github.com/CFF4HA/Staircase/internal/types"
)

func HandleDatasourceDELETE(w http.ResponseWriter, r *http.Request) error {
	var request types.DatabaseDatasource
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		if err == io.EOF {
			request = types.DatabaseDatasource{}
		} else {
			core.Logger.Error("error decoding request body", "error", err)
			return err
		}
	}
	db := database.Database()

	user, err := getUserIdFromSession(r)
	if err != nil {
		core.Logger.Error("error fetching user id from session", "error", err)
		return err
	}

	tx := db.Model(&types.DatabaseDatasource{}).Where("owner_id = ?", user).Where("id = ?", request.ID).Delete(&types.DatabaseDatasource{})
	if tx.Error != nil {
		core.Logger.Error("error deleting datasource", "error", tx.Error)
		return tx.Error
	}

	// we now delete the metadata that is tied to that
	// data source as well.
	if tx := db.Model(&types.DatabaseDatasourceMetadata{}).Where("datasource_id = ?", request.ID).Delete(&types.DatabaseDatasourceMetadata{}); tx.Error != nil {
		core.Logger.Error("error deleting datasource metadata", "error", tx.Error)
		return tx.Error
	}

	core.Logger.Debug("deleted datasource with id", "id", request.ID)
	return nil
}

func HandleDatasourceGET(w http.ResponseWriter, r *http.Request) error {
	var request types.DatabaseDatasource
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		if err == io.EOF {
			request = types.DatabaseDatasource{}
		} else {
			core.Logger.Error("error decoding request body", "error", err)
			return err
		}
	}
	db := database.Database()

	user, err := getUserIdFromSession(r)
	if err != nil {
		core.Logger.Error("error fetching user id from session", "error", err)
		return err
	}

	var response any
	if request.ID != uuid.Nil {
		ds := &types.DatabaseDatasource{}
		if tx := db.Where("owner_id = ?", user).Preload("Metadata").First(&ds, "id = ?", request.ID); tx.Error != nil {
			core.Logger.Error("error fetching datasource", "error", tx.Error)
			return tx.Error
		}

		response = ds
	} else {
		var dss []types.DatabaseDatasource
		if tx := db.Where("owner_id = ?", user).Where(&request).Preload("Metadata").Find(&dss); tx.Error != nil {
			core.Logger.Error("error fetching datasources", "error", tx.Error)
			return tx.Error
		}

		response = dss
	}

	return json.NewEncoder(w).Encode(response)
}

func HandleDatasourcePUT(w http.ResponseWriter, r *http.Request) error {
	var request types.DatabaseDatasource
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		core.Logger.Error("error decoding request body", "error", err)
		return err
	}

	user, err := getUserIdFromSession(r)
	if err != nil {
		core.Logger.Error("error fetching user id from session", "error", err)
		return err
	}

	uid := uuid.New()

	db := database.Database()
	ds := &types.DatabaseDatasource{
		ID:                 uid,
		Name:               request.Name,
		DatabaseConnString: request.DatabaseConnString,
		DatabaseTableName:  request.DatabaseTableName,
		DatabaseColumnName: request.DatabaseColumnName,
		Metadata:           types.DatabaseDatasourceMetadata{ID: uuid.New(), DatasourceId: uid},
		OwnerId:            user,
	}

	ds_tx := db.Model(&types.DatabaseDatasource{}).Create(ds)
	if ds_tx.Error != nil {
		core.Logger.Error("error creating datasource", "error", ds_tx.Error)
		return ds_tx.Error
	}

	return nil
}
