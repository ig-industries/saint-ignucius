package discord

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

// Session wraps a discordgo.Session type.
type Session struct {
	*discordgo.Session
}

// NewSession creates a new session, adds relevant handlers, and returns a
// Session that can be opened.
func NewSession(token string) (*Session, error) {
	bot, err := discordgo.New(fmt.Sprintf("Bot %s", token))
	if err != nil {
		return nil, fmt.Errorf("[discord] %s", err)
	}

	bot.AddHandler(messageCreate)

	return &Session{bot}, nil
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	author := m.Author

	if author.ID == s.State.User.ID {
		return
	}

	user := &User{
		ID:            author.ID,
		Username:      author.Username,
		Discriminator: author.Discriminator,
		Avatar:        author.Avatar,
		Bot:           author.Bot,
		MFAEnabled:    author.MFAEnabled,
		Locale:        author.Locale,
		Verified:      author.Verified,
		Email:         author.Email,
	}

	fmt.Println(user)

	_, err := s.ChannelMessageSend(m.ChannelID, m.Content)
	if err != nil {
		log.Fatalf("[discord] %s", err)
	}
}
