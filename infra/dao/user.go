package dao

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/uekiGityuto/go-practice/domain/entity"
	"golang.org/x/xerrors"
)

type User struct {
	db *sql.DB
}

func NewUser(db *sql.DB) *User {
	return &User{
		db: db,
	}
}

func (dao *User) Find(id uuid.UUID) (*entity.User, error) {
	row := dao.db.QueryRow("SELECT * FROM user WHERE id = ?", id.String())
	if row.Err() != nil {
		return nil, xerrors.Errorf("idによるユーザ情報の取得に失敗しました。: %w", row.Err())
	}
	var user entity.User
	if err := row.Scan(&user.ID, &user.FamilyName, &user.GivenName, &user.Age, &user.Sex); err != nil {
		// 存在しない場合にここに入る。
		// TODO: 現状バリデーションエラー以外がシステムエラーになってしまっているので直したい。
		return nil, xerrors.Errorf("'id=%s' のユーザ情報は存在しません。: %w", id.String(), err)
	}
	return &user, nil
}

func (dao *User) Save(user *entity.User) error {
	const sql = "INSERT INTO user (id, family_name, given_name, age, sex) VALUES (?, ?, ?, ?, ?)"
	_, err := dao.db.Exec(sql, user.ID, user.FamilyName, user.GivenName, user.Age, user.Sex)
	if err != nil {
		return xerrors.Errorf("ユーザ情報のDBへの登録に失敗しました。: %w", err)
	}
	return nil
}
