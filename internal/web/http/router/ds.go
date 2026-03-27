package router

import (
	"html/template"
	"net/http"

	"github.com/CFF4HA/Staircase/internal/types"
	"github.com/CFF4HA/Staircase/internal/web/core"
	"github.com/google/uuid"
)

var (
	tmplDatasource *template.Template
)

func HandleDatasourceGET(w http.ResponseWriter, r *http.Request) error {
	id := r.FormValue("id")
	_ = id

	model := &types.DatabaseDatasource{
		ID:                 uuid.New(),
		Name:               "Example Datasource",
		DatabaseConnString: "postgresql://user:password@localhost:5432/mydb",
		DatabaseTableName:  "my_table",
		DatabaseColumnName: "my_column",
	}

	// this will ensure the template is initialized and
	// ready to be used
	if tmplDatasource == nil {
		t, err := template.ParseFiles(
			core.TemplateDir + "/datasource/ds.html",
		)
		if err != nil {
			core.Logger.Error("failed to parse template", "error", err)
		}

		tmplDatasource = t
	}

	w.Header().Set("HX-Retarget", "#modal-content-target")
	w.Header().Set("HX-Trigger", "showModal")

	return tmplDatasource.Execute(w, model)
}
