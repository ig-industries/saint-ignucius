package discord

import (
	"encoding/json"
	"log"
)

// User represents a Discord user, by the API definition of a user. This is
// basically a clone of the discordgo.User but for right now it's just included
// for testing stuff out. This will be expanded later to include other things
// for some other services.
type User struct {
	ID            string `json:"id,omitempty"`
	Username      string `json:"username,omitempty"`
	Discriminator string `json:"discriminator,omitempty"`
	Avatar        string `json:"avatar,omitempty"`
	Bot           bool   `json:"bot,omitempty"`
	MFAEnabled    bool   `json:"mfa_enabled,omitempty"`
	Locale        string `json:"locale,omitempty"`
	Verified      bool   `json:"verified,omitempty"`
	Email         string `json:"email,omitempty"`
	Flags         int    `json:"flags,omitempty"`
	PremiumType   int    `json:"premium_type,omitempty"`
}

// UserRepository is a collection of methods that provide DB interaction. As long as
// these methods are implemented, anyone can use any type of DB service they want.
type UserRepository interface {
	UserExists(string) (bool, error)
	SelectUserByID(string) (*User, error)
	InsertUser(*User) error
}

// String implements the stringer interface. For now, the struct is just
// marshalled to a JSON format, but in the future we might want to do something
// like `username#discriminator, points` once the trivia service is up.
func (user User) String() string {
	userBytes, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		log.Fatalf("[discord] %s", err)
	}

	return string(userBytes)
}
