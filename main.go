package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {

	go httpClient()
	// http server

	log.Print("starting server...")
	http.HandleFunc("/", handler)

	// Determine port for HTTP service.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}

	// Start HTTP server.
	log.Printf("listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("NAME")
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(w, "Hello %s!\n", name)
}

func httpClient() {
	url := "http://34.149.109.170/"

	for i := 0; i < 1000; i++ {

		req, _ := http.NewRequest("GET", url, nil)

		client := new(http.Client)
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error Request:", err)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != 200 {
			fmt.Println("Error Response:", resp.Status)
			return
		}

		byteArray, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(byteArray))
	}
}
