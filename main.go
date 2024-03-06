package main

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

var templates = template.Must(template.New("").ParseGlob("*.html"))

// TODO: Read flags from env vars to keep them static
var flag_1 = generateFlag()
var flag_2 = generateFlag()
var flag_3 = generateFlag()

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/cookies", cookiesHandler)
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		encodedFlag := base64.StdEncoding.EncodeToString([]byte(flag_1))

		data := map[string]string{
			"EncodedFlag": encodedFlag,
		}

		err := templates.ExecuteTemplate(w, "index.html", data)
		if err != nil {
			http.Error(w, "Server error", http.StatusInternalServerError)
		}

	case "TRACE":
		w.Header().Set("Ctf", flag_2)
		w.WriteHeader(http.StatusOK)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func cookiesHandler(w http.ResponseWriter, r *http.Request) {
	cookieResponses := [20]string{
		"Chocolate cookie",
		"Apple cookie",
		"Vanilla cookie",
		"Strawberry cookie",
		"Lemon cookie",
		"Orange cookie",
		"Special cookie",
		"Banana cookie",
		"Blueberry cookie",
		"Raspberry cookie",
		"Mint cookie",
		"Coconut cookie",
		"Pineapple cookie",
		"Mango cookie",
		"Peach cookie",
		"Grape cookie",
		"Kiwi cookie",
		"Lime cookie",
		"Cherry cookie",
		"Blackberry cookie",
	}

	switch r.Method {
	case "GET":
		err := templates.ExecuteTemplate(w, "cookies.html", nil)
		if err != nil {
			http.Error(w, "Server error", http.StatusInternalServerError)
		}

	case "POST":
		message := "Not a special cookie"

		if cookie, err := r.Cookie("name"); err == nil {
			// Convert the `name` cookie value to an integer index
			var index int
			fmt.Sscan(cookie.Value, &index)

			// Check if the index is within the bounds of the array
			if index >= 0 && index < len(cookieResponses) {
				if index == 6 {
					http.SetCookie(w, &http.Cookie{Name: "name", Value: flag_3, Path: "/"})
					message = cookieResponses[index]
				} else {
					http.SetCookie(w, &http.Cookie{Name: "name", Value: strconv.Itoa(index), Path: "/"})
					message = cookieResponses[index]
				}
			} else {
				http.SetCookie(w, &http.Cookie{Name: "name", Value: "-1", Path: "/"})
			}
		} else {
			http.SetCookie(w, &http.Cookie{Name: "name", Value: "-1", Path: "/"})
		}

		data := map[string]string{
			"Message": message,
		}

		err := templates.ExecuteTemplate(w, "cookies.html", data)
		if err != nil {
			http.Error(w, "Server error", http.StatusInternalServerError)
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func generateFlag() string {
	bytes := make([]byte, 16)
	_, _ = rand.Read(bytes)
	return "CTF_" + hex.EncodeToString(bytes)
}
