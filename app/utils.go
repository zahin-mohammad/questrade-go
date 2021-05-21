package app

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

func PrettyJSON(obj interface{}) string {
	json, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		log.Fatalf(err.Error())
	}
	return string(json)
}

// https://stackoverflow.com/a/37533144
func intSliceToString(
	delim string,
	a ...int,
) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}
