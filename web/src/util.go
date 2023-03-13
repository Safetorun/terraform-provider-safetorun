package main

import (
	"encoding/json"
	"log"
)

func ToJson(obj any) string {
	result, err := json.Marshal(obj)
	if err != nil {
		log.Fatal(err)
		return "Sorry. Something went wrong."
	}

	return string(result)
}
