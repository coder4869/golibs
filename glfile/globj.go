package glfile

import (
	"encoding/json"
	"io/ioutil"
)

func FillObjWithJsonFile(obj interface{}, filePath string) (err error) {
	var content []byte

	if content, err = ioutil.ReadFile(filePath); err != nil {
		return err
	}

	return json.Unmarshal(content, obj)
}
