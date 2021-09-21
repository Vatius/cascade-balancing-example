package main

import (
	"encoding/json"
	"flag"
	"fmt"
	model "github.com/Vatius/cascade-balancing-example"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// Balancer ...
type Balancer struct {
	maxRequests     int
	currentRequests int
	slaveUrl        string
}

func (b *Balancer) init(limit int, slaveUrl string) {
	b.maxRequests = limit
	b.slaveUrl = slaveUrl
	ticker := time.NewTicker(time.Second)
	go func() {
		for range ticker.C {
			b.currentRequests = 0
		}
	}()
}

func (b *Balancer) handleWithLimit(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		b.currentRequests++
		if b.currentRequests > b.maxRequests {
			if b.slaveUrl == "" {
				log.Println("limit requests")
				w.WriteHeader(http.StatusServiceUnavailable)
				return
			}
			resp, err := http.Post(b.slaveUrl, "text/plain", r.Body)
			if err != nil {
				log.Println("cant use slave server:", err)
				w.WriteHeader(http.StatusServiceUnavailable)
				return
			}
			defer resp.Body.Close()
			if resp.StatusCode == 200 {
				_, err = fmt.Fprintf(w, "{ \"success\": true}")
				if err != nil {
					log.Println(err)
					w.WriteHeader(http.StatusServiceUnavailable)
				}
				return
			} else {
				log.Println("slave server unavailable")
				w.WriteHeader(http.StatusServiceUnavailable)
				return
			}
		}
		handler(w, r)
	}
}

var (
	bindAddress = flag.String("bind", "localhost:8080", "use -bind=127.0.0.1:8081")
	slaveAddr   = flag.String("slave", "", "use -slave=127.0.0.1:8082")
	maxRequests = flag.Int("max", 3, "use -max=3")
)

func main() {
	flag.Parse()
	log.Println("Starting server on", *bindAddress)
	log.Println("Maximum requests:", *maxRequests)
	if *slaveAddr != "" {
		log.Println("Slave server:", *slaveAddr)
	} else {
		log.Println("Without slave server")
	}
	balancer := new(Balancer)
	balancer.init(*maxRequests, *slaveAddr)
	http.HandleFunc("/", balancer.handleWithLimit(handlerPostPayload))
	log.Fatal(http.ListenAndServe(*bindAddress, nil))
}

func handlerPostPayload(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var records []model.Payload
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("cant read body", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(buf, &records)
	if err != nil {
		log.Println("cant unmarshal json", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	printRecords(records...)
	_, err = fmt.Fprintf(w, "{ \"success\": true}")
	if err != nil {
		log.Println("cant write response", err)
	}
}

func printRecords(rec ...model.Payload) {
	log.Println("Received data:")
	for _, item := range rec {
		log.Println(item)
	}
}
