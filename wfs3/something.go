package wfs3

import (
	"bytes"
	"html/template"
	"io/ioutil"
)

var pageTemplate string = `
<html>
<head>
  Content-Type: text/html
  Title = {{index . "title"}}
</head>
<body>
  {{index . "content"}}
</body>
</html>
`

var rootTemplate string = `  <ul>
  {{range $index, $element := .Links}}
    <li>{{.MarshalHTML}}</li>
  {{end}}
  </ul>
`

var linkTemplate string = `{{.Title}} - ({{.Rel}}/{{.Type}}) {{.Href}} {{.Hreflang}}`

func DemoRootContentPageCreation(r *RootContent, title string) string {
	var rslt bytes.Buffer
	rc := r.MarshalHTML()
	t, err := template.New("rootPage").Parse(pageTemplate)
	if err != nil {
		panic("Couldn't parse pageTemplate")
	}
	data := map[string]interface{}{"title": title, "content": template.HTML(rc)}
	if err := t.Execute(&rslt, data); err != nil {
		panic("DemoRootContentPageCreation can't Execute() template")
	}
	brslt, err := ioutil.ReadAll(&rslt)
	return string(brslt)
}
