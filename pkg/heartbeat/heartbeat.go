package heartbeat

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
)

const NotAvailableMessage = "Not available"

var CommitHash string
var StartTime time.Time

type Message struct {
	Status string `json:"status"`
	Build  string `json:"build"`
	Uptime string `json:"uptime"`
}

func init() {
	StartTime = time.Now()
}

func handler(rw http.ResponseWriter, r *http.Request) {
	hash := os.Getenv("DEVOPS_GIT_SHA")
	if hash == "" {
		hash = NotAvailableMessage
	}
	uptime := time.Since(StartTime).String()
	err := json.NewEncoder(rw).Encode(Message{"running", hash, uptime})
	if err != nil {
		log.Fatalf("Failed to write heartbeat message. Reason: %s", err.Error())
	}
}

func Router() (http.Handler, error) {
	r := chi.NewRouter()

	r.Get("/", handler)
	return r, nil
}
