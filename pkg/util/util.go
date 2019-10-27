package util

import "encoding/json"

// Stringify :make struc to string json
func Stringify(data interface{}) string {
	dataByte, _ := json.Marshal(data)
	return string(dataByte)
}
