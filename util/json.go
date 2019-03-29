package util

import (
	"encoding/json"
	"fmt"
)

func PrettyMap(data map[string]interface{}) string {
	d, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("Pretty map err: ", err)
		return ""
	}
	return string(d)
}
