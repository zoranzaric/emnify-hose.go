package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"net/http"
)

func getJsonString() []byte {
	if rand.Intn(1000)%1000 == 0 {
		return []byte(`{"id": "bar", "event_type": {"id": 6}}`)
	} else {
		return []byte(`{"id": "bar", "event_type": {"id": 1}}`)
	}
}

func main() {
	var c int = 0
	for {
		response, err := http.Post("http://127.0.0.1:6666/", "application/json", bytes.NewBuffer(getJsonString()))

		if err != nil {
			panic(err)
		}

		c += 1

		if c%1000 == 0 {
			fmt.Println(c)
		}

		defer response.Body.Close()
	}
}
