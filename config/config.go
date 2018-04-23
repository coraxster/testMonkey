package config

import "encoding/json"

type Config struct {
	Bind string `json:"bind"`
	Endpoints []Endpoint `json:"endpoints"`
}

type Endpoint struct {
	Uri string `json:"uri"`
	Method string `json:"method"`
	Status int `json:"status"`
	Response json.RawMessage
}

