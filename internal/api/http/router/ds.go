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

	tx := db.Model(&types.DatabaseDatasource{}).Where("id = ?", request.ID).Delete(&types.DatabaseDatasource{})
	if tx.Error != nil {
		core.Logger.Error("error deleting datasource", "error", tx.Error)
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

	var response any
	if request.ID != uuid.Nil {
		ds := &types.DatabaseDatasource{}
		if tx := db.First(&ds, "id = ?", request.ID); tx.Error != nil {
			core.Logger.Error("error fetching datasource", "error", tx.Error)
			return tx.Error
		}

		response = ds
	} else {
		var dss []types.DatabaseDatasource
		if tx := db.Where(&request).Find(&dss); tx.Error != nil {
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

	uuid := uuid.New()

	db := database.Database()
	ds := &types.DatabaseDatasource{
		ID:                 uuid,
		Name:               request.Name,
		DatabaseConnString: request.DatabaseConnString,
		DatabaseTableName:  request.DatabaseTableName,
		DatabaseColumnName: request.DatabaseColumnName,
	}

	if tx := db.Model(&types.DatabaseDatasource{}).Create(ds); tx.Error != nil {
		core.Logger.Error("error creating datasource", "error", tx.Error)
		return tx.Error
	}

	return nil
}
