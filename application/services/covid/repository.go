package covid

import (
	"github.com/thedevsaddam/gojsonq/v2"
)

//FindOne ...
func (c *CovidService) FindOne(filter *Filter, db *gojsonq.JSONQ) (*map[string]interface{}, error) {
	result := map[string]interface{}{}

	return &result, nil
}

//Find ...
func (c *CovidService) Find(filter *Filter, db *gojsonq.JSONQ) (*[]map[string]interface{}, int64, error) {
	result := []map[string]interface{}{}

	return &result, int64(len(result)), nil
}

func _BuildFilter(filter *Filter) ([]map[string]interface{}, []map[string]interface{}) {
	queryAnd := []map[string]interface{}{}
	queryOr := []map[string]interface{}{}

	// if filter.Status != nil {
	// 	queryAnd = append(queryAnd, map[string]interface{}{"status": *filter.Status})
	// }

	// if len(queryAnd) > 0 {
	// 	for _, each := range queryAnd {
	// 		tx.Where(each)
	// 	}
	// }

	// if len(queryOr) > 0 {
	// 	for _, each := range queryOr {
	// 		if len(queryOr) > 1 {
	// 			tx.Or(each)
	// 		} else {
	// 			tx.Where(each)
	// 		}
	// 	}
	// }

	return queryAnd, queryOr
}
