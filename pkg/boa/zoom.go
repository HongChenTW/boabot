package boa

import (
	"net/http"
)

// SlackResponser returns BoA response in Slack message format
func ZoomResponser(r *http.Request) (interface{}, error) {
	switch r.Method {
	case http.MethodGet, http.MethodPost:
		// allowed methods
	default:
		return nil, Error(http.StatusMethodNotAllowed)
	}

	return nil, nil
}
