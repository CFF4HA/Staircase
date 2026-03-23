package router

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"

	"github.com/CFF4HA/Staircase/internal/api/core"
	"github.com/CFF4HA/Staircase/internal/api/database"
	"github.com/CFF4HA/Staircase/internal/types"
)

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
