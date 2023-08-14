package applemaps

import (
	"encoding/json"
	"io"
)

type errorResponse struct {
	Error struct {
		Message string `json:"message"`
		Details []any  `json:"details"`
	} `json:"error"`
}

func unmarshalErrorResponse(reader io.Reader) *errorResponse {
	var errorRes = &errorResponse{}
	_ = json.NewDecoder(reader).Decode(errorRes)
	return errorRes
}
