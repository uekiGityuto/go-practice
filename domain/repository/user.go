package repository

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/uekiGityuto/go-practice/domain/entity"
)

type User interface {
	Find(*sql.Tx, context.Context, uuid.UUID) (*entity.User, error)
	Save(*sql.Tx, context.Context, *entity.User) error
}
