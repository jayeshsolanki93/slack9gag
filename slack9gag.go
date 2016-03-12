package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Response struct {
	Text        string        `json:"text"`
	Attachments []Attachments `json:"attachments"`
}

type Attachments struct {
	Title      string `json:"title"`
	Title_link string `json:"title_link"`
	Thumb_url  string `json:"thumb_url"`
}

type jsonData struct {
	Status  int64
	Message string
	Data    []Data
}

type Data struct {
	Id      string
	Caption string
	Images  struct {
		Small  string
		Cover  string
		Normal string
		Large  string
	}
	Media interface{}
	Link  string
	Votes struct {
		Count int64
	}
	Comments struct {
		Count int64
	}
}

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
		r, err := http.Get("http://infinigag.k3min.eu")
		if err != nil {
			fmt.Println("Error requesting data")
			return
		}
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println("Error while opening file", err)
			return
		}
		x := new(jsonData)
		err = json.Unmarshal(body, &x)
		if err != nil {
			fmt.Println("Error while parsing file", err)
			return
		}
		jsonResp(w, x)
	} else {
		fmt.Fprint(w, "I do not understand your command.")
	}
}

func jsonResp(w http.ResponseWriter, x *jsonData) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	attachments := make([]Attachments, len(x.Data))
	for i := 0; i < len(x.Data); i++ {
		attachments[i] = Attachments{
			Title:      x.Data[i].Caption,
			Title_link: x.Data[i].Link,
			Thumb_url:  x.Data[i].Images.Small,
		}
	}

	resp := Response{
		Text:        "lorem ipsum",
		Attachments: attachments,
	}

	r, err := json.Marshal(resp)
	if err != nil {
		fmt.Println("Couldn't marshal hook response:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(r)
}
