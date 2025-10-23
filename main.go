package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"runtime"
	"time"
)

type DebugInfo struct {
	Timestamp      string            `json:"timestamp"`
	Method         string            `json:"method"`
	Path           string            `json:"path"`
	RemoteAddr     string            `json:"remote_addr"`
	Headers        map[string]string `json:"headers"`
	QueryParams    map[string]string `json:"query_params"`
	Host           string            `json:"host"`
	ServerInfo     ServerInfo        `json:"server_info"`
}

type ServerInfo struct {
	Hostname       string `json:"hostname"`
	GoVersion      string `json:"go_version"`
	NumCPU         int    `json:"num_cpu"`
	NumGoroutine   int    `json:"num_goroutine"`
	OS             string `json:"os"`
	Architecture   string `json:"architecture"`
}

func debugHandler(w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()
	
	headers := make(map[string]string)
	for name, values := range r.Header {
		if len(values) > 0 {
			headers[name] = values[0]
		}
	}
	
	queryParams := make(map[string]string)
	for key, values := range r.URL.Query() {
		if len(values) > 0 {
			queryParams[key] = values[0]
		}
	}
	
	debugInfo := DebugInfo{
		Timestamp:   time.Now().UTC().Format(time.RFC3339),
		Method:      r.Method,
		Path:        r.URL.Path,
		RemoteAddr:  r.RemoteAddr,
		Headers:     headers,
		QueryParams: queryParams,
		Host:        r.Host,
		ServerInfo: ServerInfo{
			Hostname:     hostname,
			GoVersion:    runtime.Version(),
			NumCPU:       runtime.NumCPU(),
			NumGoroutine: runtime.NumGoroutine(),
			OS:           runtime.GOOS,
			Architecture: runtime.GOARCH,
		},
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	
	json.NewEncoder(w).Encode(debugInfo)
}

func main() {
	http.HandleFunc("/", debugHandler)
	
	port := "8080"
	log.Printf("Starting server on port %s", port)
	
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}