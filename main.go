package main

import (
	//"fmt"
	"log"
	"net/http"
	"os"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hi00000ijbvgchfcfgcgfgxgx"))
}

func about(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("about"))
}

func main() {
	port := getEnvOrDefault("PORT", "8080")
	log.Println("Listening on port", port)
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)

	http.ListenAndServe(":"+port, nil)
}

func getEnvOrDefault(key, defaultValue string) string {
	result := defaultValue
	val, ok := os.LookupEnv(key)
	if ok {
		result = val
	}
	return result
}
