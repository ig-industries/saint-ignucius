package sqlite

import (
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/ig-industries/saint-ignucius/discord"
)

type userRepository struct {
	*sqlx.DB
}

var errNotImplemented = errors.New("[sqlite] not implemented")

// NewUserRepository creates the user table if it doesn't exist, otherwise it
// just returns an instance of userRepository which wraps an underlying sqlx.DB
func NewUserRepository(db *sqlx.DB) (discord.UserRepository, error) {
	// create user table if it doesn't exist
	return &userRepository{db}, nil
}

func (ur *userRepository) UserExists(ID string) (bool, error) {
	var exists bool

	err := ur.QueryRow("SELECT EXISTS(SELECT * FROM users WHERE id = ? LIMIT 1)", ID).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("[sqlite] %s", err)
	}

	return exists, nil
}

func (ur *userRepository) SelectUserByID(ID string) (*discord.User, error) {
	user := &discord.User{}
	err := ur.Get(&user, "SELECT * FROM users WHERE id = ?", ID)
	if err != nil {
		return nil, fmt.Errorf("[sqlite] %s", err)
	}

	return user, nil
}

func (ur *userRepository) InsertUser(user *discord.User) error {
	return errNotImplemented
}
