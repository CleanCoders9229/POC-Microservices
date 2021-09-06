package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func main() {
	msg := `http error status: 400; body: {
		"error": {
		  "code": 400,
		  "message": "EMAIL_EXISTS",
		  "errors": [
			{
			  "message": "EMAIL_EXISTS",
			  "domain": "global",
			  "reason": "invalid"
			}
		  ]
		}
	  }`

	// sep := strings.Split(msg, ";")
	sep := strings.SplitAfterN(msg, `"error":`, 2)
	errMsg := strings.Trim(sep[1], "{}")
	var errJson map[string]interface{}
	json.Unmarshal([]byte(errMsg), &errJson)

	fmt.Println(errJson)
	fmt.Println(errJson["code"])
	errDetail := errJson["errors"].([]interface{})[0]
	fmt.Println(errDetail)
	fmt.Println(errDetail.(map[string]interface{})["message"])
}
