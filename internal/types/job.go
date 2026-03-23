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

	ID uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;not null"`

	Name         string              `json:"name"`
	Staircases   []DatabaseStaircase `json:"staircases" gorm:"foreignKey:JobId;constraint:OnDelete:CASCADE;"`
	Metadata     DatabaseJobMetadata `json:"metadata" gorm:"foreignKey:JobId;constraint:OnDelete:CASCADE;"`
	DatasourceId uuid.UUID           `json:"datasource_id" gorm:"type:uuid;not null;"`
}

type DatabaseStaircase struct {
	gorm.Model

	ID          uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;not null"`
	JobId       uuid.UUID `json:"job_id" gorm:"type:uuid;not null"`
	Declaration string    `json:"declaration" gorm:"not null"`
}

type DatabaseJobMetadata struct {
	gorm.Model

	ID            uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;not null"`
	JobId         uuid.UUID `json:"job_id" gorm:"type:uuid;not null"`
	IsInitialized bool      `json:"is_initialized" gorm:"default:false"`
	IsInProgress  bool      `json:"is_in_progress" gorm:"default:false"`
	IsFinished    bool      `json:"is_finished" gorm:"default:false"`
	IsUploaded    bool      `json:"is_uploaded" gorm:"default:false"`
	LastScan      time.Time `json:"last_scan" gorm:"default:null"`
}

type DatabaseDatasource struct {
	gorm.Model

	ID                 uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;not null"`
	Name               string    `json:"name"`
	DatabaseConnString string    `json:"database_conn_string" gorm:"not null"`
	DatabaseTableName  string    `json:"database_table_name" gorm:"not null"`
	DatabaseColumnName string    `json:"database_column_name" gorm:"not null"`

	Jobs []DatabaseJob `json:"jobs" gorm:"foreignKey:DatasourceId;constraint:OnDelete:CASCADE;"`
}
