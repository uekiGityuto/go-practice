package entity

import "github.com/google/uuid"

type User struct {
	ID         uuid.UUID
	FamilyName string
	GivenName  string
	Age        int
	Sex        string
}

func NewUser(familyName string, givenName string, age int, sex string) (*User, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	user := &User{
		ID:         id,
		FamilyName: familyName,
		GivenName:  givenName,
		Age:        age,
		Sex:        sex,
	}
	return user, nil
}
