package main

import (
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

const servicePrefix = `http://127.0.0.1:1992`

func TestCompare(t *testing.T) {
	go listenAndServe(":1992")
	time.Sleep(time.Second * 1)

	res, err := http.Get(servicePrefix + "/compare?user=testdude111&user=rowantestacc")
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	buf, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("buf: %v", string(buf))

	if res.StatusCode != 200 {
		t.Fatalf("Expecting 200")
	}

}
