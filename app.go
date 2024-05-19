package main

import (
	"encoding/base64"
	"fmt"
	"regexp"
	"strings"

	"github.com/atotto/clipboard"
)

func main() {
	contents,err := clipboard.ReadAll()
	if err != nil {
		fmt.Println("Error reading clipboard contents")
	}
	jwtRegex := regexp.MustCompile(`([a-zA-Z0-9\-_]+)\.([a-zA-Z0-9\-_]+)\.([a-zA-Z0-9\-_]+)`)
	matches := jwtRegex.FindAllStringSubmatch(contents, -1)
	if len(matches) > 0 {
		fmt.Println("Found JWT Token")
		header := matches[0][1]
		decodedHeader, err := decodeBase64URL(header)
		if err != nil {
			fmt.Println("Error decoding header")
			return 
		}
		payload := matches[0][2]
		fmt.Println(payload)
		decodedPayload, err := decodeBase64URL(payload)
		if err != nil {
			fmt.Println("Error decoding payload")
			return
		}
		signature := matches[0][3]
		decodedSignature, err := decodeBase64URL(signature)
		if err != nil {
			fmt.Println("Error decoding signature")
			return
		}
		printedResult := fmt.Sprintf("Header: %s\nPayload: %s\nSignature: %s\n", string(decodedHeader), string(decodedPayload), string(decodedSignature))
		clipboard.WriteAll(printedResult)
		fmt.Println(printedResult)
		return
	}
	decoded, err := decodeBase64URL(contents)
	if err != nil {
		return
	}
	printedResult := string(decoded)
	fmt.Println(printedResult)
	clipboard.WriteAll(printedResult)

}

func decodeBase64URL(encoded string) ([]byte, error) {
	encoded = encoded + strings.Repeat("=", (4-len(encoded)%4)%4)
	decoded, err := base64.URLEncoding.DecodeString(encoded)
	if err != nil {
		return nil, err
	}
	return decoded, nil
}
