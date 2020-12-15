package gofocused

import (
	"encoding/base64"
	"fmt"
)

func EncodeDecodeStr() {
	var data = "Yohan Soon"

	var encodedStr = base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println("Encoded:", encodedStr)

	var decodedByte, _ = base64.StdEncoding.DecodeString(encodedStr)
	var decodedStr = string(decodedByte)
	fmt.Println("Decoded:", decodedStr)
}

func EncodeDecode() {
	var data = "Yohan Soon"

	var encoded = make([]byte, base64.StdEncoding.EncodedLen(len(data)))
	base64.StdEncoding.Encode(encoded, []byte(data))
	var encodedString = string(encoded)
	fmt.Println(encodedString)

	var decoded = make([]byte, base64.StdEncoding.DecodedLen(len(encoded)))
	var _, err = base64.StdEncoding.Decode(decoded, encoded)
	if err != nil {
		fmt.Println(err.Error())
	}
	var decodedString = string(decoded)
	fmt.Println(decodedString)
}
