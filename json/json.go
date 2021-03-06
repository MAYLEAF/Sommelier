/*
Package Json implement a simple library for json CRUD.
*/
package json

import (
	"encoding/json"
	"github.com/MAYLEAF/Sommelier/logger"
	"io"
	"io/ioutil"
	"os"
	"reflect"
	"regexp"
	"strings"
)

type Json struct {
	json map[string]interface{}
}

func Decode(r io.Reader, v map[string]interface{}) error {
	dec := json.NewDecoder(r)
	if err := dec.Decode(&v); err != nil {
		logger.Info("Json Decode error occur Err: %v, Interface: %v", err, v)
		return err
	}
	return nil
}

func Encode(w io.Writer, v interface{}) error {
	enc := json.NewEncoder(w)
	if err := enc.Encode(v); err != nil {
		logger.Info("Json Encode error occur Err: %v, Interface: %v", err, v)
		return err
	}
	return nil
}

func Modify(node interface{}, k string, v interface{}) interface{} {
	js := make(map[string]interface{})
	var err error
	var byteJson []byte
	if byteJson, err = json.Marshal(node); err != nil {
		panic(err)
	}
	if json.Unmarshal(byteJson, &js); err != nil {
		panic(err)
	}
	js[k] = v
	return js
}

func Read(v interface{}) []byte {
	msg, err := json.Marshal(v)
	if err != nil {
		logger.Error("Fail to read json err:%v \n\n", err)
		return nil
	}
	return msg
}

func ReadJsonFile(file_name string, v interface{}) interface{} {
	var File *os.File
	var err error
	if File, err = os.Open(file_name); err != nil {
		panic(err)
	}
	defer File.Close()

	byteValue, _ := ioutil.ReadAll(File)
	mType := reflect.TypeOf(v).Elem()
	newV := reflect.New(mType).Interface()
	json.Unmarshal([]byte(byteValue), newV)
	return newV
}

func (e *Json) Json() map[string]interface{} {
	return e.json
}

func (e *Json) SetJson(json map[string]interface{}) {
	e.json = json
}

func (e *Json) Load(key string) interface{} {
	if value, ok := e.json[key]; ok {
		return value
	}
	return nil
}

func (e *Json) Contains(key string, value string) bool {
	if e.json[key] == nil {
		return false
	}
	re := regexp.MustCompile(`(.*)` + value + `(.*)`)
	msg, err := json.Marshal(e.json[key])
	if err != nil {
		logger.Info("Fail to read json err: %v \n\n", err)
	}
	if re.Find(msg) == nil {
		return false
	}
	return true
}
func (e *Json) Select(key string) *Json {
	var msg []byte
	var err error
	var newJson = Json{}
	if msg, err = json.Marshal(e.json[key]); err != nil {
		logger.Info("Fail to read json err: %v \n", err)
	}
	dec := json.NewDecoder(strings.NewReader(string(msg)))
	if err = dec.Decode(&newJson.json); err != nil {
		logger.Info("Json Select Error: %v", err)
	}
	return &newJson
}
