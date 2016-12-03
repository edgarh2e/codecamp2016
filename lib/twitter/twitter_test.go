package twitter

import "testing"

func TestGetUser(t *testing.T) {
	u, err := GetUser("xiam")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("u: %v", u)

	followers, err := GetFollowers("xiam")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("followers: %v", len(followers))

	following, err := GetFollowing("xiam")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("following: %v", len(following))
}
