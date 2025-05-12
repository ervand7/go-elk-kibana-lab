package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	msg := fmt.Sprintf(
		`{"timestamp":"%s","level":"INFO","msg":"Request received","method":"%s","path":"%s"}`,
		time.Now().Format(time.RFC3339),
		r.Method,
		r.URL.Path,
	)
	fmt.Fprintln(w, "OK")
	log.Println(msg)
}

func main() {
	// Log file setup
	logFile, err := os.OpenFile("/app/logs/app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	log.SetOutput(logFile)

	http.HandleFunc("/", handler)
	log.Println(`{"timestamp":"` + time.Now().Format(time.RFC3339) + `","level":"INFO","msg":"Starting server on :8080"}`)
	http.ListenAndServe(":8080", nil)
}
