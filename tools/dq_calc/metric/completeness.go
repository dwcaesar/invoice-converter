package metric

import (
	"encoding/json"
	"log"
	"reflect"
	"strings"
)

func IsAddressComplete(rec interface{}) bool {
	fields := [4]string{"City", "Zip", "Street", "Name"}

	return IsSimpleRecordComplete(fields[:], rec)
}

func IsItemsComplete(rec interface{}) bool {
	items, isConverted := rec.([]interface{})
	if !isConverted {
		log.Panicf("cannot convert %s", reflect.TypeOf(rec))
		return false
	}

	fields := [4]string{"Name", "Amount", "Vat", "ItemPrice"}
	for _, item := range items {
		if !IsSimpleRecordComplete(fields[:], item) {
			return false
		}
	}

	return true
}
func IsSimpleRecordComplete(requiredFields []string, record interface{}) bool {
	var result map[string]string
	jsonBytes, err := json.Marshal(record)
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(jsonBytes, &result); err != nil {
		panic(err)
	}

	for _, f := range requiredFields {
		if v, exist := result[f]; exist {
			str := strings.TrimSpace(v)
			if len(str) == 0 {
				log.Default().Printf("the record's field %s was incomplete", f)
				return false
			}
		} else {
			log.Default().Printf("the record's field %s not existing", f)
			return false
		}
	}

	return true
}
