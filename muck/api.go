package muck

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/bwmarrin/discordgo"
)

const contentType = "application/json"

type muckAPI struct {
	host   string
	token  string
	client *http.Client
}

type muckMessage struct {
	GuildID   string `json:"guild_id"`
	ChannelID string `json:"channel_id"`
	UserID    string `json:"user_id"`
	Content   string `json:"content"`
	Edit      bool   `json:"is_edit"`
}

func newMessage(m *discordgo.Message, g string, e bool) *muckMessage {
	return &muckMessage{
		GuildID:   g,
		ChannelID: m.ChannelID,
		UserID:    m.Author.ID,
		Content:   m.Content,
		Edit:      e,
	}
}

func (api *muckAPI) url() string {
	return fmt.Sprintf("%s/api/muck", api.host)
}

func (api *muckAPI) newRequest(r io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodPost, api.url(), r)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", api.token)
	req.Header.Set("Content-Type", contentType)

	return req, nil
}

func (api *muckAPI) sendMessage(m *muckMessage) {
	payload, err := json.Marshal(m)
	if err != nil {
		report(err)
		return
	}
	r := bytes.NewReader(payload)

	req, err := api.newRequest(r)
	if err != nil {
		report(err)
		return
	}

	// TODO: handle response
	_, err = api.client.Do(req)
	if err != nil {
		report(err)
		return
	}
	fmt.Printf("(muck) sent data <edit: %v>\n", m.Edit)
}
