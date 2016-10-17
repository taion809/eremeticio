package main

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/pressly/chi"
	"github.com/pressly/chi/middleware"
)

var (
	httpListenAddr   string = ":5555"
	httpEremeticAddr string
)

func init() {
	if listenAddr := os.Getenv("ENV_HTTP_ADDR"); listenAddr != "" {
		httpListenAddr = listenAddr
	}

	httpEremeticAddr = os.Getenv("ENV_EREMETIC_ADDR")
}

func main() {
	client := NewClient(httpEremeticAddr)

	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.CloseNotify)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		taskList, err := client.FetchTasks(r.Context())
		if err != nil {
			writeError(w, err)
			return
		}

		if err = writeJson(w, taskList, http.StatusOK); err != nil {
			writeError(w, err)
			return
		}
	})

	http.ListenAndServe(httpListenAddr, r)
}

func writeError(w http.ResponseWriter, value error) error {
	return writeJson(w, struct{ errors []string }{errors: []string{value.Error()}}, http.StatusInternalServerError)
}

func writeJson(w http.ResponseWriter, value interface{}, status int) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(value); err != nil {
		return err
	}

	return nil
}
