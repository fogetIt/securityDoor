package utils

import (
	"encoding/json"
)


func CreateMessage(status int, msg string, extMaps ...map[string]interface{}) string {
	var mapObj = make(map[string]interface{}, 3)
	mapObj["status"] = status
	mapObj["msg"] = msg
	for i := range extMaps {
		for extKey := range extMaps[i] {
			mapObj[extKey] = extMaps[i][extKey]
		}
	}

	byteArray, err := json.Marshal(mapObj)
	if err != nil {
		return ""
	}
	return string(byteArray[:])
}
