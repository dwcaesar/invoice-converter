package metric

import (
	"daquam/report"
	"encoding/json"
	"log"
	"reflect"
	"strings"
)

func IsAddressComplete(rec interface{}, parent string, report *report.Report) bool {
	fields := [4]string{"City", "Zip", "Street", "Name"}

	return IsSimpleRecordComplete(fields[:], rec, parent, report)
}

func IsItemsComplete(rec interface{}, parent string, csvReport *report.Report) bool {
	items, isConverted := rec.([]interface{})
	if !isConverted {
		log.Panicf("cannot convert %s", reflect.TypeOf(rec))
		return false
	}

	fields := [4]string{"Name", "Amount", "Vat", "ItemPrice"}
	for _, item := range items {
		if !IsSimpleRecordComplete(fields[:], item, parent, csvReport) {
			return false
		}
	}

	return true
}
func IsSimpleRecordComplete(requiredFields []string, record interface{}, parent string, r *report.Report) bool {
	var result map[string]string
	jsonBytes, err := json.Marshal(record)
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(jsonBytes, &result); err != nil {
		panic(err)
	}

	for _, f := range requiredFields {
		r.NewEntry(parent+"/"+f, "IsSimpleRecordComplete")
		if v, exist := result[f]; exist {
			str := strings.TrimSpace(v)
			if len(str) == 0 {
				log.Default().Printf("the record's field %s was incomplete", f)
				r.CloseEntry(str, false, "field is empty")
				return false
			} else {
				r.CloseEntry(str, true, "")
			}
		} else {
			log.Default().Printf("the record's field %s not existing", f)
			r.CloseEntry(nil, false, "field is missing")
			return false
		}
	}

	return true
}
