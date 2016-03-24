package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type Response struct {
	Text        string        `json:"text"`
	Attachments []Attachments `json:"attachments"`
}

type Attachments struct {
	Title      string `json:"title"`
	Title_link string `json:"title_link"`
	Image_url  string `json:"image_url"`
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

	url := "http://infinigag.k3min.eu"
	// Read the Request Parameter "command"
	command := r.FormValue("command")

	if command == "/9gag" {
		// Read the Request Parameter "text"
		text := r.FormValue("text")
		s := strings.Split(text, " ")
		var section string
		var subsection string
		if len(s) == 2 {
			section = s[0]
			subsection = s[1]
		} else if len(s) == 1 {
			section = s[0]
		}
		switch section {
		case "":
		case "cute":
			url += "/cute"
		case "comic":
			url += "/comic"
		case "cosplay":
			url += "/cosplay"
		case "design":
			url += "/design"
		case "food":
			url += "/food"
		case "funny":
			url += "/funny"
		case "geeky":
			url += "/geeky"
		case "gif":
			url += "/gif"
		case "girl":
			url += "/girl"
		case "meme":
			url += "/meme"
		case "nsfw":
			url += "/nsfw"
		case "timely":
			url += "/timely"
		case "wtf":
			url += "/wtf"
		default:
			fmt.Fprint(w, "I do not understand your command.")
			return
		}

		switch subsection {
		case "":
		case "fresh":
			url += "/fresh"
		case "hot":
			url += "/hot"
		default:
			fmt.Fprint(w, "I do not understand your command.")
			return
		}

		fmt.Println(url)
		r, err := http.Get(url)
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
			Image_url:  x.Data[i].Images.Small,
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
