package models

import (
	"context"
	"github.com/WytQuant/bunframework/connectdb"
	"time"
)

type User struct {
	ID        int       `json:"id" bun:"id,pk,autoincrement"`
	FirstName string    `json:"first_name" bun:"first_name,notnull"`
	LastName  string    `json:"last_name" bun:"last_name,notnull"`
	Email     string    `json:"email" bun:"email,notnull"`
	Password  string    `json:"password" bun:",notnull"`
	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
}

func CreateUser(user *User) error {
	_, err := connectdb.Database.Db.NewInsert().
		Model(user).
		Exec(context.Background())
	return err
}

func GetUser(id int) (*User, error) {
	var user User

	err := connectdb.Database.Db.
		NewSelect().Model(&user).Where("id = ?", id).
		Scan(context.Background())
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func CheckEmail(email string, user *User) bool {
	err := connectdb.Database.Db.NewSelect().
		Model(user).
		Column("id", "first_name", "last_name", "email", "password").
		Where("email = ?", email).
		Limit(1).
		Scan(context.Background())
	if err != nil {
		return false
	}

	return true
}
