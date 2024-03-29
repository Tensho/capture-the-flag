// TODO: Check 12 factors in app
package ctf

import (
	"embed"
	"encoding/base64"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

//go:embed index.html cookies.html
var content embed.FS

var templates = template.Must(template.New("").ParseFS(content, "index.html", "cookies.html"))

var flag_1 = os.Getenv("FLAG_1")
var flag_2 = os.Getenv("FLAG_2")
var flag_3 = os.Getenv("FLAG_3")

func init() {
	functions.HTTP("ctf", EntryPoint)
}

func EntryPoint(w http.ResponseWriter, r *http.Request) {
	// Route based on the URL path
	switch r.URL.Path {
	case "/":
		indexHandler(w, r)
	case "/cookies":
		cookiesHandler(w, r)
	default:
		http.NotFound(w, r)
	}
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

	case "HEAD":
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
