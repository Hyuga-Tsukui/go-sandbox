package main

import (
	"context"
	"log"
	"net/http"
	"server/worker"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	queue := worker.RunWorker(ctx, 10)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	http.HandleFunc("/task-a", func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Recovered in executeTask: %v", r)
			}
		}()
		switch r.Method {
		case http.MethodPost:
			queue <- &worker.Task{ID: 1}
			w.WriteHeader(http.StatusAccepted)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/shutdown-worker", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			cancel()
			w.WriteHeader(http.StatusAccepted)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
