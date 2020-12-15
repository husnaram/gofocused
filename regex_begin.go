package gofocused

import (
	"fmt"
	"regexp"
)

var text = "banana burger soup"
var regex, err = regexp.Compile(`[a-z]+`)

func RegexBegin() {
	var res1 = regex.FindAllString(text, 2)
	fmt.Printf("%#v \n", res1)

	var res2 = regex.FindAllString(text, -1)
	fmt.Printf("%#v \n", res2)
}

func UseMathStr() {
	var isMatch = regex.MatchString(text)
	fmt.Println(isMatch)
}
