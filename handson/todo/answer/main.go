package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"syscall"
)

type TODO struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	User string `json:"user"`
	Done bool   `json:"done"`
}

var (
	mu    = sync.RWMutex{}
	cache = make(map[int]TODO)
	maxID = 1
)

func Get(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(strings.TrimPrefix(r.RequestURI, "/"))
	if err != nil || id == 0 {
		err = json.NewEncoder(w).Encode(cache)
		if err != nil {
			log.Println(err)
		}
		return
	}
	mu.RLock()
	todo, ok := cache[id]
	mu.RUnlock()
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(todo)
	if err != nil {
		log.Println(err)
	}
}

func Create(w http.ResponseWriter, r *http.Request) {
	var todo TODO
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if todo.ID == 0 {
		todo.ID = maxID
		maxID++
	}
	mu.Lock()
	cache[todo.ID] = todo
	mu.Unlock()
	w.WriteHeader(http.StatusAccepted)
	err = json.NewEncoder(w).Encode(todo)
	if err != nil {
		log.Println(err)
	}
}

func Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(strings.TrimPrefix(r.RequestURI, "/"))
	if err != nil || id == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	mu.RLock()
	todo, ok := cache[id]
	mu.RUnlock()
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	todo.Done = true
	mu.Lock()
	cache[id] = todo
	mu.Unlock()
	w.WriteHeader(http.StatusAccepted)
	err = json.NewEncoder(w).Encode(todo)
	if err != nil {
		log.Println(err)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(strings.TrimPrefix(r.RequestURI, "/"))
	if err != nil || id == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	mu.Lock()
	delete(cache, id)
	mu.Unlock()
	w.WriteHeader(http.StatusAccepted)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	} else if !strings.HasPrefix(port, ":") && len(port) < 5 {
		port = ":" + port
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method)
		switch r.Method {
		case http.MethodGet:
			Get(w, r)
		case http.MethodPost:
			Create(w, r)
		case http.MethodPut, http.MethodPatch:
			Update(w, r)
		case http.MethodDelete:
			Delete(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	srv := &http.Server{
		Addr:    port,
		Handler: mux,
	}
	srv.SetKeepAlivesEnabled(true)

	ech := make(chan error, 1)
	sigCh := make(chan os.Signal, 1)
	defer func() {
		close(sigCh)
		close(ech)
	}()

	go func() {
		ech <- srv.ListenAndServe()
	}()

	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT)

	for {
		select {
		case <-sigCh:
			srv.Shutdown(context.Background())
			return
		case err := <-ech:
			log.Fatalln(err)
		}
	}
}
