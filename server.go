package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

var startedAt = time.Now()

func main() {
	http.HandleFunc("/config-map-fruits", ConfigMapFruits)
	http.HandleFunc("/healthz", Healthz)
	http.HandleFunc("/secret", Secret)
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":8000", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("NAME")
	age := os.Getenv("AGE")
	fmt.Fprintf(w, "Hello, I'm %s and I'm %s years old", name, age)
}

func ConfigMapFruits(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("/go/config/fruits.txt")
	if err != nil {
		log.Fatal("Error reading file: ", err)
	}
	fmt.Fprintf(w, "Config Map Fruits: %s", string(data))
}

func Secret(w http.ResponseWriter, r *http.Request) {
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	fmt.Fprintf(w, "User: %s, Password: %s", user, password)
}

func Healthz(w http.ResponseWriter, r *http.Request) {
	duration := time.Since(startedAt)
	if duration.Seconds() < 10 {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Duration: %s", duration)
	} else {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "OK: %s", duration)
	}
}
