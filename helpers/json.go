package helpers

import (
	"encoding/json"
	"errors"
	"log"
)

// GetJSON ...
func GetJSON(doc interface{}) (map[string]interface{}, error) {
	var jsonDoc map[string]interface{}

	b, err := json.Marshal(doc)
	if err != nil {
		log.Println(err)
		return nil, errors.New("Error al covertir a JSON")
	}

	if err := json.Unmarshal(b, &jsonDoc); err != nil {
		return nil, errors.New("Error al covertir a JSON")
	}

	return jsonDoc, nil
}

// MergeMap ...
func MergeMap(destination, source map[string]interface{}) map[string]interface{} {
	if destination == nil || source == nil {
		return map[string]interface{}{}
	}

	for key, value := range source {
		if v, isOk := value.(map[string]interface{}); isOk {
			d := map[string]interface{}{}

			if destinationValue, isOk := destination[key].(map[string]interface{}); isOk {
				d = destinationValue
			}

			destination[key] = MergeMap(d, v)
		} else {
			destination[key] = value
		}
	}

	return destination
}
