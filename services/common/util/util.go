package util

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ParseJSON(r *http.Request, v any) error{
	if r.Body == nil {
		return fmt.Errorf("missing req body")
	}

	return json.NewDecoder(r.Body).Decode(v)
}
