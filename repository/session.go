package repository

import (
	"database/sql"
	"errors"
	"forwardchaining/model"
	"time"
)

type SessionsRepository interface {
	AddSessions(session model.Session) error
	DeleteSession(token string) error
	UpdateSessions(session model.Session) error
	SessionAvailName(name string) error
	SessionAvailToken(token string) (model.Session, error)

	FetchByID(id int) (*model.Session, error)
}

type sessionsRepoImpl struct {
	db *sql.DB
}

func NewSessionRepo(db *sql.DB) *sessionsRepoImpl {
	return &sessionsRepoImpl{db}
}

func (u *sessionsRepoImpl) AddSessions(session model.Session) error {
	_, err := u.db.Exec(`INSERT INTO sessions (token, username, expiry) VALUES ($1, $2, $3)`, session.Token, session.Username, session.Expiry)
	if err != nil {
		return err
	}

	return nil // TODO: replace this
}

func (u *sessionsRepoImpl) DeleteSession(token string) error {
	_, err := u.db.Exec("DELETE FROM sessions WHERE token = $1", token)
	if err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (u *sessionsRepoImpl) UpdateSessions(session model.Session) error {
	_, err := u.db.Exec("UPDATE sessions SET token = $1, expiry = $2 WHERE username = $3", session.Token, session.Expiry, session.Username)
	if err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (u *sessionsRepoImpl) SessionAvailName(name string) error {
	var session model.Session

	row := u.db.QueryRow("SELECT id, username, token FROM sessions WHERE username = $1", name)
	err := row.Scan(&session.ID, &session.Username, &session.Token)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errors.New("session not found")
		}
		return err
	}

	return nil // TODO: replace this
}

func (u *sessionsRepoImpl) SessionAvailToken(token string) (model.Session, error) {
	var session model.Session
	query := "SELECT id, username, token, expiry FROM sessions WHERE token = $1"
	err := u.db.QueryRow(query, token).Scan(&session.ID, &session.Username, &session.Token, &session.Expiry)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.Session{}, errors.New("session not found")
		}
		return model.Session{}, err
	}

	currentTime := time.Now().UTC()
	if session.Expiry.Before(currentTime) {
		return model.Session{}, errors.New("Token is Expired!")
	}

	return session, nil // TODO: replace this
}

func (u *sessionsRepoImpl) FetchByID(id int) (*model.Session, error) {
	row := u.db.QueryRow("SELECT id, token, username, expiry FROM sessions WHERE id = $1", id)

	var session model.Session
	err := row.Scan(&session.ID, &session.Token, &session.Username, &session.Expiry)
	if err != nil {
		return nil, err
	}

	return &session, nil
}
