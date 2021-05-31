package database

import gojsonq "github.com/thedevsaddam/gojsonq/v2"

// Init ...
/*
Initializes de Database
*/

func Init(databaseFileURL string) *gojsonq.JSONQ {
	return gojsonq.New().File(databaseFileURL)
}
