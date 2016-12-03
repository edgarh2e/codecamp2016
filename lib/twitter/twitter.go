package twitter

import (
	"errors"
	"fmt"
	//"log"
	"github.com/ChimeraCoder/anaconda"
	"github.com/boltdb/bolt"
	"github.com/edgarh2e/codecamp2016/data"

	"encoding/json"
	"log"
	"net/url"
	"time"
)

var api *anaconda.TwitterApi

var db *bolt.DB

var (
	usersBucket = []byte("users")
)

func init() {
	anaconda.SetConsumerKey("rzOULAw1RmqCSbhH9Mwlas6kF")
	anaconda.SetConsumerSecret("SvJRKbWO1d3A6VzRMqZEgbR67aLtF3Lp3Vo94RWfd9LbluuzRz")
	api = anaconda.NewTwitterApi("4830555837-VYIXrgP3ehP4wtWgQfsqGFS6yvLlAsMwjSxqeV7", "H2JohUEU4q65SlglPCMws7aPk1Vx2pFEih254heTEKttb")

	var err error
	db, err = bolt.Open("my.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}

	buckets := [][]byte{usersBucket}

	db.Update(func(tx *bolt.Tx) error {
		for _, bucket := range buckets {
			_, err := tx.CreateBucketIfNotExists(bucket)
			if err != nil {
				log.Fatal(err)
			}
		}
		return nil
	})

}

func cacheUsers(users ...data.User) error {
	tx, err := db.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	bucket := tx.Bucket([]byte(usersBucket))
	for _, user := range users {
		data, err := json.Marshal(user)
		if err != nil {
			return err
		}
		bucket.Put([]byte(user.ID), data)
	}

	return tx.Commit()
}

func getUserFromCache(id int64) (*data.User, error) {
	var user data.User
	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(usersBucket))
		v := bucket.Get([]byte(fmt.Sprintf("%d", id)))
		if v == nil {
			return errors.New("Not in cache")
		}
		if err := json.Unmarshal(v, &user); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUser(username string) (*data.User, error) {
	u, err := api.GetUsersShow(username, nil)
	if err != nil {
		return nil, err
	}
	return &data.User{
		ID:       u.IdStr,
		Username: u.ScreenName,
	}, nil
}

func GetFollowing(username string) ([]data.User, error) {
	pages := api.GetFriendsIdsAll(url.Values{"screen_name": {username}})

	ids := []int64{}
	for page := range pages {
		ids = append(ids, page.Ids...)
	}

	return getUsersByIds(ids)
}

func GetFollowers(username string) ([]data.User, error) {
	pages := api.GetFollowersIdsAll(url.Values{"screen_name": {username}})

	ids := []int64{}
	for page := range pages {
		ids = append(ids, page.Ids...)
	}

	return getUsersByIds(ids)
}

func getUsersByIds(ids []int64) ([]data.User, error) {
	users := make([]data.User, 0, len(ids))

	var missing []int64
	for _, lookUpId := range ids {
		if cached, err := getUserFromCache(lookUpId); err == nil {
			users = append(users, *cached)
		} else {
			missing = append(missing, lookUpId)
		}
	}

	log.Printf("unknopwn: %v", len(missing))

	var step, max = 100, len(missing)

	for i := 0; true; i++ {
		a := i * step
		if a >= max {
			break
		}

		b := (i + 1) * step
		if b > max {
			b = max
		}

		us, err := api.GetUsersLookupByIds(missing[a:b], nil)
		if err != nil {
			return nil, err
		}
		for _, u := range us {
			users = append(users, data.User{
				ID:       u.IdStr,
				Username: u.ScreenName,
			})
		}
	}

	cacheUsers(users...)

	return users, nil
}
