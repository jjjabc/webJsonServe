package models

import "encoding/json"

const (
	FAILED  = "failed"
	SUCCESS = "success"
)

type RespJSON struct {
	Status string
	Des    string
	Data   *json.RawMessage
}
