package glnet

import (
	//	"fmt"
	"net/http"
	"strings"
)

func FormParse(w http.ResponseWriter, r *http.Request) map[string]string {
	r.ParseForm() // parameters analysis, default not analysed

	var paramMap map[string]string = make(map[string]string, 1)
	for k, v := range r.Form {
		paramMap[k] = strings.Join(v, "")
		//		fmt.Println("key:", k)
		//		fmt.Println("val:", strings.Join(v, ""))
	}

	//	fmt.Println("r.form", r.Form)
	//	fmt.Println("path", r.URL.Path)
	//	fmt.Println("scheme", r.URL.Scheme)
	//	fmt.Println("url_long", r.Form["url_long"])

	return paramMap
}
