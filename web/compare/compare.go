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

	tgf := `1 uno
2 dos
3 tres
4 cuatro
5 cinco
6 seis
7 siete
#
1 2
3 4
1 5
2 5
6 6
7 1
7 2
7 3
2 7
6 1`

	cmd := exec.Command("python", binDir+"tgf2svg.py")
	cmd.Stdin = bytes.NewBuffer([]byte(tgf))

	stdOut, stdErr := bytes.NewBuffer(nil), bytes.NewBuffer(nil)

	err := cmd.Run()
	if err != nil {
		return nil, err
	}

	log.Printf("out: %v", string(stdOut.Bytes()))
	log.Printf("err: %v", string(stdErr.Bytes()))

	hash := fmt.Sprintf("%x", md5.Sum(stdOut.Bytes()))

	fileName := path.Join(dataDir, hash+".svg")

	fp, err := os.Open(fileName)
	if err == nil {
		// TODO: is it a dir?
		return &Output{
			Image: path.Base(fileName),
			Data:  data,
		}, nil
	}

	fp, err = os.Create(fileName)
	log.Printf("filename: %v", fileName)
	if err != nil {
		return nil, err
	}
	defer fp.Close()

	if _, err := fp.Write(stdOut.Bytes()); err != nil {
		return nil, err
	}

	return &Output{
		Image: path.Base(fileName),
		Data:  data,
	}, nil
}
