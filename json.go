package tool

import (
	"encoding/json"
	"fmt"
	"time"
)

type JSONTime time.Time

const (
	timeFormat = "2006-01-02 15:04:05"
)

func (t *JSONTime) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+timeFormat+`"`, string(data), time.Local)
	*t = JSONTime(now)
	return
}

func (t JSONTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(timeFormat)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, timeFormat)
	b = append(b, '"')
	return b, nil
}

func (t JSONTime) String() string {
	return time.Time(t).Format(timeFormat)
}

func Dump(entity ...interface{}) {
	switch entity[0].(type) {
	case string:
		StringDump(entity...)
	default:
		JSONDump(entity...)
	}
}

func JSONDump(entities ...interface{}) {

	var dumps []string

	for _, entity := range entities {

		dump, err := json.MarshalIndent(&entity, "", "\t")

		if err != nil {
			panic("JSONDump:" + err.Error())
		}

		dumps = append(dumps, string(dump))

	}

	for _, dump := range dumps {
		_, err := fmt.Println(dump)

		if err != nil {
			panic("JSONDump:" + err.Error())
		}
	}
}

func StringDump(entities ...interface{}) {

	for _, entity := range entities {
		_, err := fmt.Println(entity)

		if err != nil {
			panic("StringDump:" + err.Error())
		}
	}
}
