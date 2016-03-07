package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("listening...")
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Read the Request Parameter "command"
	command := r.FormValue("command")

	if command == "/9gag" {
		fmt.Fprint(w, "Hello World")
	} else {
		fmt.Fprint(w, "I do not understand your command.")
	}
}