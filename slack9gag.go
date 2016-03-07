package slack9gag

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/", handler)
	fmt.Println("listening...")
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