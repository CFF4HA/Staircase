package types

import (
	"github.com/google/uuid"

	"time"

	"gorm.io/gorm"
)

// We will use this meta structure to collect jobs from the
// database context.
type DatabaseJob struct {
	gorm.Model

	ID uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;not null;default:gen_random_uuid()"`

	Name         string              `json:"name"`
	Frequency    int                 `json:"frequency" gorm:"not null;default:60;"`
	Staircases   []DatabaseStaircase `json:"staircases" gorm:"foreignKey:JobId;constraint:OnDelete:CASCADE;"`
	Metadata     DatabaseJobMetadata `json:"metadata" gorm:"foreignKey:JobId;constraint:OnDelete:CASCADE;"`
	DatasourceId uuid.UUID           `json:"datasource_id" gorm:"type:uuid;not null;"`
	OwnerId      uuid.UUID           `json:"owner_id" gorm:"type:uuid;not null;"`
}

type DatabaseStaircase struct {
	gorm.Model

	ID          uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;not null;default:gen_random_uuid()"`
	JobId       uuid.UUID `json:"job_id" gorm:"type:uuid;not null"`
	Declaration string    `json:"declaration" gorm:"not null"`
}

type DatabaseJobMetadata struct {
	gorm.Model

	ID            uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;not null;default:gen_random_uuid()"`
	JobId         uuid.UUID `json:"job_id" gorm:"type:uuid;not null"`
	IsInitialized bool      `json:"is_initialized" gorm:"default:false"`
	IsInProgress  bool      `json:"is_in_progress" gorm:"default:false"`
	IsFinished    bool      `json:"is_finished" gorm:"default:false"`
	IsUploaded    bool      `json:"is_uploaded" gorm:"default:false"`
	LastScan      time.Time `json:"last_scan" gorm:"default:null"`

	// the following was added to allow for saving job results
	// in the database for admin panel demonstration.
	InjectedJavascript string `json:"injected_javascript" gorm:"type:text;default:''"`
	Result             string `json:"result" gorm:"default:''"`
	Image              string `json:"image" gorm:"type:text;default:''"`
	Error              string `json:"error" gorm:"default:''"`
}

type DatabaseDatasource struct {
	gorm.Model

	ID                 uuid.UUID                  `json:"id" gorm:"type:uuid;primaryKey;not null;default:gen_random_uuid()"`
	Name               string                     `json:"name"`
	DatabaseConnString string                     `json:"database_conn_string" gorm:"not null"`
	DatabaseTableName  string                     `json:"database_table_name" gorm:"not null"`
	DatabaseColumnName string                     `json:"database_column_name" gorm:"not null"`
	OwnerId            uuid.UUID                  `json:"owner_id" gorm:"type:uuid;not null"`
	Metadata           DatabaseDatasourceMetadata `json:"metadata" gorm:"foreignKey:DatasourceId;constraint:OnDelete:CASCADE;"`

	Jobs []DatabaseJob `json:"jobs" gorm:"foreignKey:DatasourceId;constraint:OnDelete:CASCADE;"`
}

type DatabaseDatasourceMetadata struct {
	gorm.Model

	// meta information
	ID           uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;not null;default:gen_random_uuid()"`
	DatasourceId uuid.UUID `json:"datasource_id" gorm:"type:uuid;not null"`

	// actual useful bussiness logic information
	LastPing time.Time `json:"last_ping" gorm:"default:null"`
	IsAlive  bool      `json:"is_alive" gorm:"default:false"`
	Error    string    `json:"error"`
}
