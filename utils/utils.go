package utils

import (
	"encoding/json"
)

type Model struct{}

func (u *Model) ToJson() ([]byte, error) {
	return json.Marshal(u)
}
