package main

import (
	"bytes"
	"encoding/json"
	"log"
	"math"
	"net/http"
	"os"
	"sync"
	"time"
)

type Routes struct {
	Requests []Request `json:"routes"`
}

type Request struct {
	Method  string            `json:"method"`
	Url     string            `json:"url"`
	Headers map[string]string `json:"headers"`
	Body    map[string]string `json:"body"`
}

func loadRequest(filePath string) *Routes {
	readjson, err := os.ReadFile(filePath)
	if err != nil {
		panic("Failed to read file spec")
	}
	var routes Routes
	json.Unmarshal(readjson, &routes)
	if len(routes.Requests) < 1 {
		panic("Should have at least one request")
	}
	return &routes
}

func (r *Request) fireRequest(wg *sync.WaitGroup, resp *Rapport, ch *chan struct{}) {
	now := time.Now()

	client := &http.Client{}
	body_, err := json.Marshal(r.Body)
	if err != nil {
		log.Fatal("erro body")
	}
	req, err := http.NewRequest(r.Method, r.Url, bytes.NewBuffer([]byte(body_)))
	for key, value := range r.Headers {
		req.Header.Set(key, value)
	}
	if err != nil {
		log.Fatal(err)
	}
	resp_, err := client.Do(req)
	if err != nil {
		log.Fatal()
		log.Fatal("Error request")
	}
	duration := time.Since(now).Seconds()
	rounded := math.Round(duration*1000) / 1000

	resp.mu.Lock()
	defer resp.mu.Unlock()

	url_ := r.Url + "__" + r.Method
	if _, exists := resp.mapping[url_]; !exists {
		resp.mapping[url_] = map[string][]float32{
			"times":  {},
			"status": {},
		}
	}
	resp.mapping[url_]["times"] = append(resp.mapping[url_]["times"], float32(rounded))
	resp.mapping[url_]["status"] = append(resp.mapping[url_]["status"], float32(resp_.StatusCode))
	<-*ch
	wg.Done()
}
