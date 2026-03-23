package bridges

import (
	"net/http"
)

type BridgeBackend struct {
	Backend *string
}

func (b BridgeBackend) Data(w http.ResponseWriter, r *http.Request) (any, error) {
	return b.Backend, nil
}
