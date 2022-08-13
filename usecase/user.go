package usecase

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/uekiGityuto/go-practice/domain/entity"
	"github.com/uekiGityuto/go-practice/domain/repository"
	"golang.org/x/xerrors"
)

type User struct {
	repo repository.User
	db   *sql.DB
}

func NewUser(repo repository.User, db *sql.DB) *User {
	return &User{
		repo: repo,
		db:   db,
	}
}

func (uc *User) Find(id string) (*entity.User, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, xerrors.Errorf("uuidのパースに失敗しました。: %w", err)
	}

	tx, err := uc.db.Begin()
	if err != nil {
		return nil, xerrors.Errorf("トランザクション開始が失敗しました。: %w", err)
	}
	user, err := uc.repo.Find(tx, uid)
	if err != nil {
		tx.Rollback()
		return nil, xerrors.Errorf("ユーザ情報の取得に失敗しました。: %w", err)
	} else if user == nil {
		tx.Rollback()
		return nil, xerrors.Errorf("ユーザ情報の取得に失敗しました。: %w", NotFoundErr)
	}
	if err := tx.Commit(); err != nil {
		return nil, xerrors.Errorf("トランザクションのコミットが失敗しました。: %w", err)
	}

	return user, nil
}

func (uc *User) Save(familyName string, givenName string, age int, sex string) (*uuid.UUID, error) {
	user, err := entity.NewUser(familyName, givenName, age, sex)
	if err != nil {
		return nil, xerrors.Errorf("userエンティティの生成に失敗しました。")
	}

	tx, err := uc.db.Begin()
	if err != nil {
		return nil, xerrors.Errorf("トランザクション開始が失敗しました。: %w", err)
	}
	err = uc.repo.Save(tx, user)
	if err != nil {
		tx.Rollback()
		return nil, xerrors.Errorf("userエンティティの登録に失敗しました。")
	}
	if err := tx.Commit(); err != nil {
		return nil, xerrors.Errorf("トランザクションのコミットが失敗しました。: %w", err)
	}

	return &user.ID, nil
}
