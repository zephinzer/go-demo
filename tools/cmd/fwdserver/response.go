package main

import (
	"encoding/json"
)

type Response struct {
	Data interface{} `json:"data"`
	Error string `json:"error"`
	NextHop *NextHop `json:"nextHop"`
}

func (r Response) Serialize() string {
	return string(r.ToBytes())
}

// ToBytes is for use with the http.ResponseWriter.Write([]byte) function
func (r Response) ToBytes() []byte {
	marshalledResponse, marshalErr := json.Marshal(r)
	if marshalErr != nil {
		return Response{
			Error: marshalErr.Error(),
		}.ToBytes()
	}
	return marshalledResponse
}
