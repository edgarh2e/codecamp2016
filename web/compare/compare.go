package compare

import (
	"bytes"
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/edgarh2e/codecamp2016/data"
	"github.com/edgarh2e/codecamp2016/lib/twitter"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
)

const dataDir = "/tmp"

const binDir = "./../bin/"

type Output struct {
	Image string      `json:"image"`
	Data  []data.User `json:"data"`
}

func readImage(file string) ([]byte, error) {
	fileName := path.Base(path.Clean(file))
	fp, err := os.Open(path.Join(dataDir, fileName))
	if err != nil {
		return nil, err
	}
	defer fp.Close()

	return ioutil.ReadAll(fp)
}

func compare(usernames ...string) (*Output, error) {
	if len(usernames) == 0 {
		return nil, errors.New("Missing input")
	}

	data := make([]data.User, 0, 2)

	for _, username := range usernames {
		user, err := twitter.GetUser(username)
		if err != nil {
			return nil, err
		}

		user.Followers, err = twitter.GetFollowers(username)
		if err != nil {
			return nil, err
		}

		user.Following, err = twitter.GetFollowing(username)
		if err != nil {
			return nil, err
		}

		data = append(data, *user)
	}

	users := make(map[string]string)
	nodes := map[string]map[string]bool{}

	for _, user := range data {
		users[user.ID] = user.Username
		for _, f := range user.Followers {
			users[f.ID] = f.Username
			if nodes[f.ID] == nil {
				nodes[f.ID] = make(map[string]bool)
			}
			nodes[f.ID][user.ID] = true
		}
		for _, f := range user.Following {
			users[f.ID] = f.Username
			if nodes[user.ID] == nil {
				nodes[user.ID] = make(map[string]bool)
			}
			nodes[user.ID][f.ID] = true
		}
	}

	log.Printf("users: %v", users)
	log.Printf("nodes: %v", nodes)

	filterUsers := make(map[string]string)

	tgfNodes := ""
	for a, aa := range nodes {
		for b := range aa {
			if nodes[a][b] && nodes[b][a] {
				filterUsers[a] = users[a]
				filterUsers[b] = users[b]
				tgfNodes += fmt.Sprintf("%s %s\n", a, b)
			}
		}
	}

	if len(filterUsers) < 1 {
		return nil, errors.New("No friends")
	}

	tgf := ""
	for id, username := range filterUsers {
		tgf += fmt.Sprintf("%s %s\n", id, username)
	}
	tgf += "#\n" + tgfNodes

	log.Printf("tgf: %v", tgf)

	cmd := exec.Command("python", binDir+"tgf2svg.py")
	cmd.Stdin = bytes.NewBuffer([]byte(tgf))

	stdOut, stdErr := bytes.NewBuffer(nil), bytes.NewBuffer(nil)
	cmd.Stdout = stdOut
	cmd.Stderr = stdErr

	err := cmd.Run()
	if err != nil {
		return nil, err
	}

	buf := stdOut.Bytes()

	hash := fmt.Sprintf("%x", md5.Sum(buf))

	fileName := path.Join(dataDir, hash+".svg")

	fp, err := os.Open(fileName)
	if err == nil {
		// TODO: is it a dir?
		return &Output{
			Image: "/compare/view/" + path.Base(fileName),
			Data:  data,
		}, nil
	}

	fp, err = os.Create(fileName)
	log.Printf("filename: %v", fileName)
	if err != nil {
		return nil, err
	}
	defer fp.Close()

	if _, err := fp.Write(buf); err != nil {
		return nil, err
	}

	return &Output{
		Image: "/compare/view/" + path.Base(fileName),
		Data:  data,
	}, nil
}
