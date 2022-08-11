package repository

import (
	"github.com/google/uuid"
	"github.com/uekiGityuto/go-practice/domain/entity"
)

type User interface {
	Find(uuid.UUID) (*entity.User, error)
	Save(user *entity.User) error
}
