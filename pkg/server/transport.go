package server

import (
	"context"
	"encoding/json"
	"net/http"
)

type getResponse struct {
	Message string `json:"message"`
	Err     string `json:"err,omitempty"`
}

type exampleRequest struct{
	Key   string `json:"Key"`
	Key2   string `json:"Key2"`
}

func DecodeExampleRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req exampleRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// Last but not least, we have the encoder for the response output
func EncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
