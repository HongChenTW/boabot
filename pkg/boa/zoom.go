package boa

import (
	"fmt"
	"net/http"
)

type ZoomCommand struct {
	Event   string   `json:"event,omitempty"`
	Payload *Payload `json:"payload,omitempty"`
}

type Payload struct {
	AccountID   string `json:"accountId,omitempty"`
	ChannelName string `json:"channelName,omitempty"`
	Cmd         string `json:"cmd,omitempty"`
	Name        string `json:"name,omitempty"`
	RobotJID    string `json:"robotJid,omitempty"`
	Timestamp   string `json:"timestamp,omitempty"`
	ToJID       string `json:"toJid,omitempty"`
	UserID      string `json:"userId,omitempty"`
	UserJID     string `json:"userJid,omitempty"`
	UserName    string `json:"userName,omitempty"`
}

// Response - chatbot reponse
type Response struct {
	RobotJID  string   `json:"robot_jid,omitempty"`
	ToJID     string   `json:"to_jid,omitempty"`
	AccountID string   `json:"account_id,omitempty"`
	Content   *Content `json:"content,omitempty"`
}

type Content struct {
	Head *Head `json:"head,omitempty"`
}

type Head struct {
	Text string `json:"text,omitempty"`
}

// ZoomResponser returns BoA response in Zoom message format
func ZoomResponser(r *http.Request) (interface{}, error) {
	switch r.Method {
	case http.MethodGet, http.MethodPost:
		// allowed methods
	default:
		return nil, Error(http.StatusMethodNotAllowed)
	}

	cmd, err := ZoomCommandParse(r)
	return cmd, nil

	if err != nil {
		return nil, Error(http.StatusBadRequest)
	}

	if cmd.Payload.Cmd == "" {
		cmd.Payload.Cmd = defaultQuestion
	}

	fmt.Println(cmd)

	resp := &Response{
		RobotJID:  cmd.Payload.RobotJID,
		ToJID:     cmd.Payload.ToJID,
		AccountID: cmd.Payload.AccountID,
		Content: &Content{
			Head: &Head{
				Text: GetAnswer(),
			},
		},
	}

	return resp, nil
}

// ZoomCommandParse will parse the request of the zoom command
func ZoomCommandParse(r *http.Request) (z ZoomCommand, err error) {
	if err = r.ParseForm(); err != nil {
		return z, err
	}

	z.Event = r.PostForm.Get("event")
	z.Payload.AccountID = r.PostForm.Get("accountId")
	z.Payload.ChannelName = r.PostForm.Get("channelName")
	z.Payload.Cmd = r.PostForm.Get("cmd")
	z.Payload.Name = r.PostForm.Get("name")
	z.Payload.RobotJID = r.PostForm.Get("robotJid")
	z.Payload.Timestamp = r.PostForm.Get("timestamp")
	z.Payload.ToJID = r.PostForm.Get("toJid")
	z.Payload.UserID = r.PostForm.Get("userId")
	z.Payload.UserJID = r.PostForm.Get("userJid")
	z.Payload.UserName = r.PostForm.Get("userName")
	return z, nil
}
