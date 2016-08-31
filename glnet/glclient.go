// Copyright (c) 2016 coder4869 ( https://github.com/coder4869/golibs ). All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

/*	glnet/glclient.go
	provides GET/POST URL request methods
*/

package glnet

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

/*
	GET Request with URLStr and URLStr != ""
*/
func GETReqWithURL(URLStr string) (string, error) {
	if strings.EqualFold(URLStr, "") {
		return "", errors.New("Parameter 'URLStr' Value Error")
	}

	clent := &http.Client{}
	req, _ := http.NewRequest("GET", URLStr, nil)

	resp, err := clent.Do(req) //send
	if err != nil {
		return "", err
	}

	data, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	return string(data), err
}

/*
	POST Request with URLStr and URLStr != ""
*/
func POSTReqWithURL(URLStr string, params map[string]string) (string, error) {

	if strings.EqualFold(URLStr, "") {
		return "", errors.New("Parameter 'URLStr' Value Error")
	}

	v := url.Values{}
	for key, val := range params {
		v.Set(key, val)
	}

	body := ioutil.NopCloser(strings.NewReader(v.Encode())) //encode form data
	req, _ := http.NewRequest("POST", URLStr, body)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	//	fmt.Printf("%+v\n", req) //review structure of sending contents

	clent := &http.Client{}
	resp, err := clent.Do(req) //send
	if err != nil {
		return "", err
	}

	data, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	return string(data), err
}
