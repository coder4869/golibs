// Copyright (c) 2016 coder4869 ( https://github.com/coder4869/golibs ). All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

/*	glfile/globj.go
 */
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
