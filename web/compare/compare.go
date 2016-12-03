package compare

import (
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/edgarh2e/codecamp2016/data"
	"github.com/edgarh2e/codecamp2016/lib/twitter"
	"os"
	"path"
)

const dataDir = "/tmp"

type Output struct {
	Image string      `json:"image"`
	Data  []data.User `json:"data"`
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

	svg := `
		<svg xmlns="http://www.w3.org/2000/svg"
				xmlns:xlink="http://www.w3.org/1999/xlink">

				<path d="M50,50
								 A30,30 0 0,1 35,20
								 L100,100
								 M110,110
								 L100,0"
							style="stroke:#660000; fill:none;"/>
		</svg>
	`

	hash := fmt.Sprintf("%x", md5.Sum([]byte(svg)))

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
	if err != nil {
		return nil, err
	}
	defer fp.Close()

	if _, err := fp.Write([]byte(svg)); err != nil {
		return nil, err
	}

	return &Output{
		Image: path.Base(fileName),
		Data:  data,
	}, nil
}
