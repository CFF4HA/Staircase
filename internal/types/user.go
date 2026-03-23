package types

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Uid          uuid.UUID `json:"uid" gorm:"type:uuid;primaryKey;not null"`
	Username     string    `json:"username" gorm:"unique;not null"`
	PasswordHash string    `json:"password_hash" gorm:"not null"`

	ServiceTokens []ServiceToken `json:"service_tokens" gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE;"`
}

type ServiceToken struct {
	UserId uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`
	Token  uuid.UUID `json:"token" gorm:"type:uuid;primaryKey;not null"`
}

type Session struct {
	Session uuid.UUID `json:"session"`
}
