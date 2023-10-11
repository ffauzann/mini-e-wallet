package model

import (
	"database/sql"
	"time"
)

type Timestamp struct {
	CreatedAt time.Time
	UpdatedAt sql.NullTime
	DeletedAt sql.NullTime
}
