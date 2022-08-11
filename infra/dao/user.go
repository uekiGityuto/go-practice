package dao

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/uekiGityuto/go-practice/domain/entity"
)

type User struct{}

func NewUser() *User {
	return &User{}
}

func (_ User) Find(uuid uuid.UUID) (*entity.User, error) {
	user := &entity.User{
		ID:         uuid,
		FamilyName: "植木",
		GivenName:  "宥登",
		Age:        28,
		Sex:        "男",
	}
	return user, nil
}

func (_ User) Save(user *entity.User) error {
	fmt.Println("DBに登録しました。")
	return nil
}
