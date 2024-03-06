package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/cookies", cookiesHandler)
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		flag_2 := "ABC"
		encodedFlag := base64.StdEncoding.EncodeToString([]byte(flag_2))

		htmlContent := "<html><head><meta name=\"Ctf\" content=\"" + encodedFlag + "\"></head><body><h1>ABC Hackathon 2024</h1></body></html>"
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(htmlContent))

	case "TRACE":
		flag_3 := "XYZ"
		w.Header().Set("Ctf", flag_3)
		w.WriteHeader(http.StatusOK)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func cookiesHandler(w http.ResponseWriter, r *http.Request) {
	flag_1 := "GHI"

	// Define an array with 21 elements (0th index will not be used)
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

	htmlStart := `<html><body>`
	htmlForm := `<form action="/cookies" method="post"><input type="text" name="name"/><input type="submit" value="Submit"/></form>`
	htmlEnd := `</body></html>`

	switch r.Method {
	case "GET":
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(htmlStart + htmlForm + htmlEnd))

	case "POST":
		if cookie, err := r.Cookie("name"); err == nil {
			// Convert the cookie value to an integer index
			var index int
			fmt.Sscan(cookie.Value, &index)

			// Check if the index is within the bounds of the array
			if index >= 0 && index < len(cookieResponses) {
				response := cookieResponses[index]
				if index == 6 {
					http.SetCookie(w, &http.Cookie{Name: "name", Value: flag_1, Path: "/"})
				} else {
					http.SetCookie(w, &http.Cookie{Name: "name", Value: cookie.Value, Path: "/"})
				}
				w.Write([]byte(htmlStart + response + htmlForm + htmlEnd))
			} else {
				http.SetCookie(w, &http.Cookie{Name: "name", Value: "-1", Path: "/"})
				w.Write([]byte(htmlStart + `Not a special cookie` + htmlForm + htmlEnd))
			}
		} else {
			http.SetCookie(w, &http.Cookie{Name: "name", Value: "-1", Path: "/"})
			w.Write([]byte(htmlStart + `Not a special cookie` + htmlForm + htmlEnd))
		}

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
