package main

import (
	"bytes"
	"encoding/json"
	"flag"
	model "github.com/Vatius/cascade-balancing-example"
	"log"
	"net/http"
	"sync"
	"time"
)

var (
	serverUrl = flag.String("url", "http://localhost:8080/", "use -url=http://127.0.0.1:8082/")
	interval  = flag.Int("interval", 1000, "use -interval=3000")
)

func main() {
	flag.Parse()
	log.Printf("Starting client, server url: %v, interval: %v ms \n", *serverUrl, *interval)
	defer log.Println("Bye!")
	ticker := time.NewTicker(time.Second) // set interval
	w := sync.WaitGroup{}
	w.Add(1)
	go func() {
		for range ticker.C {
			// do http post here
			jsonBody, _ := json.Marshal([]model.Payload{{1, 2, 3, 4, 5}})
			r := bytes.NewReader(jsonBody)
			resp, err := http.Post(*serverUrl, "text/plain", r)
			if err != nil {
				log.Println("cant post server", err)
				w.Done()
				return
			}
			log.Println("Send request: ", resp.StatusCode)
		}
	}()
	w.Wait()
}
