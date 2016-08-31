// Copyright (c) 2016 coder4869 ( https://github.com/coder4869/golibs ). All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

/*	glweb/gltmpl.go
	provides html and related operations
*/
package glweb

import (
	"html/template"
	"io"
)

type Tmpl struct {
	Templates *template.Template
}

func html(v string) template.HTML { return template.HTML(v) }

func (p *Tmpl) Load() {
	p.Templates = template.Must(template.New("").Funcs(template.FuncMap{
		"html": html,
	}).Delims("[[", "]]").ParseGlob("views/*.*"))
}

func (p *Tmpl) Execute(w io.Writer, name string, data interface{}) error {
	return p.Templates.ExecuteTemplate(w, name+".tpl", data)
}
