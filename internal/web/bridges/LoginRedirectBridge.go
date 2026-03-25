package bridges

import (
	"net/http"
)

type LoginRedirectBridge struct{}

func (b LoginRedirectBridge) Data(w http.ResponseWriter, r *http.Request) (any, error) {

	w.Header().Set("Hx-Redirect", "/")
	return nil, nil
}
