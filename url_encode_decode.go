package gofocused

import (
	"encoding/base64"
	"fmt"
)

func URLEncodeDecode() {
	var data = "https://tutorialedge.net"

	var encodedStr = base64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println("Encoded:", encodedStr)

	var decodedByte, _ = base64.URLEncoding.DecodeString(encodedStr)
	var decodedStr = string(decodedByte)
	fmt.Println("Decoded:", decodedStr)
}
