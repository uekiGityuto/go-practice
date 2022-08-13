package dao

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/uekiGityuto/go-practice/domain/entity"
	"golang.org/x/xerrors"
)

type User struct {
}

func NewUser() *User {
	return &User{}
}

func (dao *User) Find(tx *sql.Tx, id uuid.UUID) (*entity.User, error) {
	row := tx.QueryRow("SELECT * FROM user WHERE id = ?", id.String())
	if row.Err() != nil {
		return nil, xerrors.Errorf("idによるユーザ情報の取得に失敗しました。: %w", row.Err())
	}
	var user entity.User
	if err := row.Scan(&user.ID, &user.FamilyName, &user.GivenName, &user.Age, &user.Sex); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, xerrors.Errorf("idによるユーザ情報の取得に失敗しました。: %w", row.Err())
	}
	return &user, nil
}

func (dao *User) Save(tx *sql.Tx, user *entity.User) error {
	const dml = "INSERT INTO user (id, family_name, given_name, age, sex) VALUES (?, ?, ?, ?, ?)"
	_, err := tx.Exec(dml, user.ID, user.FamilyName, user.GivenName, user.Age, user.Sex)
	if err != nil {
		return xerrors.Errorf("ユーザ情報のDBへの登録に失敗しました。: %w", err)
	}
	return nil
}
