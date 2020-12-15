package gofocused

import (
	"fmt"
	"html/template"
	"net/http"
)

func RunWebServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello!")
	})

	http.HandleFunc("/index", index)

	fmt.Println("Starting Web Server @ http://localhost:9090/")
	http.ListenAndServe(":9090", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "How are you?")
}

func HandleWebTemplate() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var data = map[string]string{
			"Name":    "John Wick",
			"Message": "Happy New Year",
		}

		var t, err = template.ParseFiles("template.html")
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		t.Execute(w, data)
	})

	fmt.Println("Starting Web Server at http://localhost:9090")
	http.ListenAndServe(":9090", nil)
}
