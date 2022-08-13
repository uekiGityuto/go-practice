package repository

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/uekiGityuto/go-practice/domain/entity"
)

type User interface {
	Find(*sql.Tx, uuid.UUID) (*entity.User, error)
	Save(*sql.Tx, *entity.User) error
}
