package gofocused

import (
	"fmt"
	"net/url"
)

func URLParsing() {
	var urlString = "http://kalipare.com:80/hello?name=john wick&age=27"
	var u, err = url.Parse(urlString)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("URL: %s\n", urlString)
	fmt.Printf("Host: %s\n", u.Host)
	fmt.Printf("Path: %s\n", u.Path)

	var name = u.Query()["name"][0]
	var age = u.Query()["age"][0]
	fmt.Printf("Query: name - %s, age - %s\n", name, age)
}
