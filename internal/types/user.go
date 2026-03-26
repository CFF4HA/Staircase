package types

import (
	"github.com/google/uuid"
)

type User struct {
	Uid          uuid.UUID `json:"uid" gorm:"type:uuid;primaryKey;not null"`
	Username     string    `json:"username" gorm:"unique;not null"`
	PasswordHash string    `json:"password_hash" gorm:"not null"`

	Jobs        []DatabaseJob        `json:"jobs" gorm:"foreignKey:OwnerId;references:Uid;constraint:OnDelete:CASCADE;"`
	Datasources []DatabaseDatasource `json:"datasources" gorm:"foreignKey:OwnerId;references:Uid;constraint:OnDelete:CASCADE;"`
}

type ServiceToken struct {
	UserId uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`
	Token  uuid.UUID `json:"token" gorm:"type:uuid;primaryKey;not null"`
}

type Session struct {
	Session uuid.UUID `json:"session"`
}
