package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	TargetServer string `yaml:"target_server"`
	TargetScheme string `yaml:"target_scheme"`
}

func main() {
	configData, err := ioutil.ReadFile("app.yaml")
	if err != nil {
		log.Fatalf("failed to read config file: %v", err)
	}

	var config Config
	err = yaml.Unmarshal(configData, &config)
	if err != nil {
		log.Fatalf("failed to unmarshal config data: %v", err)
	}

	targetURL := fmt.Sprintf("%s://%s", config.TargetScheme, config.TargetServer)
	fmt.Printf("Target URL: %s\n", targetURL)

	// Create a new reverse proxy that logs requests and responses
	reverseProxy := &httputil.ReverseProxy{
		Director: func(req *http.Request) {
			req.URL.Scheme = config.TargetScheme
			req.URL.Host = config.TargetServer
			req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
			req.Header.Set("X-Origin-Host", config.TargetServer)
			log.Printf("Request: %v", req)
		},
		ModifyResponse: func(resp *http.Response) error {
			log.Printf("Response: %v", resp)
			return nil
		},
	}

	// Start the server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		reverseProxy.ServeHTTP(w, r)
	})
	log.Printf("Listening on %v://%v...", config.TargetScheme, ":8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
