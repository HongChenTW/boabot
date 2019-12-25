package boa

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/nlopes/slack"
)

// SlackResponser returns BoA response in Slack message format
func SlackResponser(r *http.Request) (interface{}, error) {
	switch r.Method {
	case http.MethodGet, http.MethodPost:
		// allowed methods
	default:
		return nil, Error(http.StatusMethodNotAllowed)
	}

	cmd, err := slack.SlashCommandParse(r)
	if err != nil {
		return nil, Error(http.StatusBadRequest)
	}

	if cmd.Text == "" {
		cmd.Text = defaultQuestion
	}

	sb := &strings.Builder{}
	if cmd.UserID != "" {
		fmt.Fprintf(sb, "<@%s>: %s\n", cmd.UserID, cmd.Text)
	} else {
		fmt.Fprintf(sb, "Someone: %s\n", cmd.Text)
	}
	fmt.Fprintf(sb, "BoA: %s", GetAnswer())

	resp := &slack.Msg{
		ResponseType: slack.ResponseTypeInChannel,
		Text:         sb.String(),
	}
	return resp, nil
}
