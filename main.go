package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
	"strconv"
	"time"
)

var myAddr string
var slaveAddr string
var maxRequests int
var currentRequests int

type Record struct {
	Price    int `json:"price"`
	Quantity int `json:"quantity"`
	Amount   int `json:"amount"`
	Object   int `json:"object"`
	Method   int `json:"method"`
}

func init() {
	flag.StringVar(&myAddr, "addr", "localhost:8081", "use -addr=127.0.0.1:8081")
	flag.StringVar(&slaveAddr, "slave", "", "use -slave=127.0.0.1:8082")
	flag.IntVar(&maxRequests, "max", 3, "use -max=3")
	flag.Parse()
}

func main() {
	log.Println("Starting server on " + myAddr)
	log.Println("Maximum requests: " + strconv.Itoa(maxRequests))

	if slaveAddr != "" {
		log.Println("Slave server: " + slaveAddr)
	} else {
		log.Println("Without slave server")
	}

	ticker := time.NewTicker(time.Second)
	go func() {
		for _ = range ticker.C {
			currentRequests = 0
		}
	}()

	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		currentRequests++
		if r.Method != "POST" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		if currentRequests > maxRequests {
			//send to slave server
			if slaveAddr == "" {
				w.WriteHeader(http.StatusServiceUnavailable)
				return
			}
			resp, err := http.Post(slaveAddr, "text/plain", r.Body)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()
			if resp.StatusCode == 200 {
				_, _ = fmt.Fprintf(w, "{ \"success\": true}")
				return
			} else {
				//slave server error
				w.WriteHeader(http.StatusServiceUnavailable)
				return
			}
		}

		decoder := json.NewDecoder(r.Body)
		var rec []Record
		err := decoder.Decode(&rec)
		if err != nil {
			panic(err)
		}
		log.Println("Received data:")
		for _, item := range rec {
			log.Println(item)
		}
		_, _ = fmt.Fprintf(w, "{ \"success\": true}")
	})

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS", "PUT"},
		AllowCredentials: true,
	})

	handler := c.Handler(router)

	log.Fatal(http.ListenAndServe(myAddr, handler))
}
