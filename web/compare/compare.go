package compare

import (
	"crypto/md5"
	"fmt"
	"os"
	"path"
)

const dataDir = "/tmp"

type Input struct {
	Usernames []string `json:"usernames"`
}

type Output struct {
	Image string `json:"image"`
}

func compare(in *Input) (*Output, error) {
	data := `
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

	hash := fmt.Sprintf("%x", md5.Sum([]byte(data)))

	fileName := path.Join(dataDir, hash+".svg")

	fp, err := os.Open(fileName)
	if err == nil {
		// TODO: is it a dir?
		return &Output{
			Image: path.Base(fileName),
		}, nil
	}

	fp, err = os.Create(fileName)
	if err != nil {
		return nil, err
	}
	defer fp.Close()

	if _, err := fp.Write([]byte(data)); err != nil {
		return nil, err
	}

	return &Output{
		Image: path.Base(fileName),
	}, nil
}
