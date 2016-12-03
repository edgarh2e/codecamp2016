package data

import (
	"testing"
)

func TestEncodeUsers(t *testing.T) {

	user1 := User{
		ID:       "0001",
		Username: "user1",
		Following: []User{
			User{
				Username: "pepe21",
				ID:       "00021",
			},
			User{
				Username: "pepe22",
				ID:       "00022",
			},
			User{
				Username: "pepe24",
				ID:       "00024",
			},
		},
		Followers: []User{
			User{
				Username: "pepe21",
				ID:       "00021",
			},
			User{
				Username: "pepe22",
				ID:       "00022",
			},
			User{
				Username: "pepe23",
				ID:       "00023",
			},
			User{
				Username: "pepe24",
				ID:       "00024",
			},
		},
	}

	user2 := User{
		ID:       "0002",
		Username: "user1",
		Followers: []User{
			User{
				Username: "pepe24",
				ID:       "00024",
			},
		},
	}

	user3 := User{
		ID:       "0003",
		Username: "user1",
		Followers: []User{
			User{
				Username: "pepe21",
				ID:       "00021",
			},
			User{
				Username: "pepe22",
				ID:       "00022",
			},
			User{
				Username: "pepe24",
				ID:       "00024",
			},
		},
	}

	out, err := generateOutput(user1, user2, user3)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%v", string(out))
}
