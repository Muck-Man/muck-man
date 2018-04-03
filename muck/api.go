package muck

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

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
	MessageID string `json:"message_id"`
	Content   string `json:"content"`
	Edited    bool   `json:"edited"`
	Timestamp int64  `json:"timestamp"`
}

func newMessage(m *discordgo.Message, g string, e bool) *muckMessage {
	var t time.Time
	var err error

	if !e {
		t, err = m.Timestamp.Parse()
	} else {
		t, err = m.EditedTimestamp.Parse()
	}
	if err != nil {
		report(err)
		return nil
	}

	return &muckMessage{
		GuildID:   g,
		ChannelID: m.ChannelID,
		UserID:    m.Author.ID,
		MessageID: m.ID,
		Content:   m.Content,
		Edited:    e,
		Timestamp: t.Unix(),
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
	resp, err := api.client.Do(req)
	if err != nil {
		report(err)
		return
	}
	if resp.StatusCode >= http.StatusBadRequest {
		report(newError(resp))
		return
	}
}

type httpError struct {
	r *http.Response
}

func newError(resp *http.Response) httpError {
	return httpError{
		r: resp,
	}
}

func (e httpError) Error() string {
	return fmt.Sprintf("%s %s - %s", e.r.Request.Method, e.r.Request.URL, e.r.Status)
}
