package main

import (
	"encoding/json"
	"fmt"
)

type RequestCode struct {
	Code string
	text string
}

func main() {
	m := RequestCode{"EEEE", "FFFF"}
	body, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))

}
