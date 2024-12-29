package engine

import (
	"encoding/json"
	"os"
)

func GetEnv(key string, defaultVal ...string) string {
	value := os.Getenv(key)

	if len(value) == 0 {
		if len(defaultVal) == 0 {
			return ""
		}
		return defaultVal[0]
	}

	return value
}

func StructToJsonBytes[T any](data T) *Result[[]byte] {
	res := NewResult[[]byte]()

	jsonData, err := json.Marshal(data)
	if err != nil {
		res.WithError(err)
		return res
	}

	res.WithPureData(jsonData)
	return res
}

func SetStructToJsonBytes[T any](data T, output *[]byte) *Result[[]byte] {
	if res := StructToJsonBytes(data); !res.IsOk() {
		return res
	} else {
		*output = res.PureData()
		return res
	}
}

func JsonBytesToStruct[T any](data []byte) *Result[T] {
	res := NewResult[T]()

	var result T
	if err := json.Unmarshal(data, &result); err != nil {
		res.WithError(err)
		return res
	}

	res.WithPureData(result)
	return res
}

func SetJsonBytesToStruct[T any](data []byte, output *T) *Result[T] {
	if res := JsonBytesToStruct[T](data); !res.IsOk() {
		return res
	} else {
		*output = res.PureData()
		return res
	}
}
