package controller

import (
	"encoding/json"
	"log"
	"net/http"
)

func ToJsonBytes(data interface{}) ([]byte, error) {
	return json.Marshal(data)
}

func ToJson(data interface{}) (string, error) {
	bytes, err := ToJsonBytes(data)
	return string(bytes), err
}

func FromJsonRequest(req *http.Request, receiver interface{}) {
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(receiver)
	if nil != err {
		log.Fatal(err)
	}
}

func AddJsonContentHeader(header http.Header) {
	header.Set("Content-Type", "application/json")
}

func WriteAsJson(data interface{}, writer http.ResponseWriter) {
	jsonData, err := ToJsonBytes(data)
	if nil != err {
		log.Fatal(err)
	}

	_, err = writer.Write(jsonData)
	if nil != err {
		log.Fatal(err)
	}
}
