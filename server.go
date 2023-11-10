package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", Hello)
	http.HandleFunc("/config-map-fruits", ConfigMapFruits)
	http.ListenAndServe(":80", nil)
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
