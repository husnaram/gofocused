package gofocused

import (
	"fmt"
	"strconv"
)

func AtoiStrConv() {
	var str = "239"
	var num, err = strconv.Atoi(str)
	if err == nil {
		fmt.Println(num)
	}
	fmt.Println(str)
}
