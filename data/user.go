package user

import (
	"encoding/json"
)

type User struct {
	ID        string `json:"id_str"`
	Username  string `json:"username"`
	Following []User `json:"following,omitempty"`
	Followers []User `json:"followers,omitempty"`
}

func generateOutput(users ...User) ([]byte, error) {
	buf, err := json.Marshal(users)
	if err != nil {
		return nil, err
	}
	return buf, nil
}
