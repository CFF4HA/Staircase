package translators

import (
	"bytes"
	"net/http"
	"strings"

	"github.com/CFF4HA/Staircase/internal/web/core"
)

func sendToBackend(endpoint string, body []byte, method string, session *http.Cookie) (*http.Response, error) {
	url := strings.TrimRight(core.BackendURL, "/") + "/" + strings.TrimLeft(endpoint, "/")

	// we make the basic request
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	// we add the sesion to the request
	if session != nil {
		req.AddCookie(session)
	}

	// we make the request to the backend server

	return http.DefaultClient.Do(req)
}
