package usecase

import (
	"github.com/google/uuid"
	"github.com/uekiGityuto/go-practice/domain/entity"
	"github.com/uekiGityuto/go-practice/domain/repository"
	"golang.org/x/xerrors"
)

type User struct {
	repo repository.User
}

func NewUser(repo repository.User) *User {
	return &User{
		repo: repo,
	}
}

func (uc *User) Find(id string) (*entity.User, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, xerrors.Errorf("uuidのパースに失敗しました: %w", err)
	}
	return uc.repo.Find(uid)
}

func (uc *User) Save(familyName string, givenName string, age int, sex string) error {
	user, err := entity.NewUser(familyName, givenName, age, sex)
	if err != nil {
		return xerrors.Errorf("userエンティティの生成に失敗しました。")
	}
	return uc.repo.Save(user)
}
