package repository

import (
	"database/sql"
	"errors"
	"forwardchaining/model"
)

type UserRepository interface {
	Add(user model.User) error
	CheckAvail(user model.User) error
	FetchByID(id int) (*model.User, error)
	CheckPass(user model.User, password string) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) UserRepository {
	return &userRepository{db}
}

func (u *userRepository) Add(user model.User) error {
	_, err := u.db.Exec(`INSERT INTO users (username, password) VALUES ($1, $2)`, user.Username, user.Password)
	if err != nil {
		return err
	}

	return nil
}

func (u *userRepository) CheckAvail(user model.User) error {
	var count int
	err := u.db.QueryRow("SELECT COUNT(*) FROM users WHERE username = $1 AND password = $2", user.Username, user.Password).Scan(&count)
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("user not found")
	}
	return nil
}

func (u *userRepository) FetchByID(id int) (*model.User, error) {
	row := u.db.QueryRow("SELECT id, username, password FROM users WHERE id = $1", id)

	var user model.User
	err := row.Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *userRepository) CheckPass(user model.User, password string) error {
	var dbPassword string
	err := u.db.QueryRow("SELECT password FROM users WHERE username = $1", user.Username).Scan(&dbPassword)
	if err != nil {
		return err
	}
	if dbPassword != password {
		return errors.New("incorrect password")
	}
	return nil
}
