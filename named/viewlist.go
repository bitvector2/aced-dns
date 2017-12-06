package named

import (
	"bytes"
	"text/template"
)

const (
	viewListTemplate = `
{{- range .Views }}
    {{ . }}
{{- end }}
`
)

type ViewList struct {
	template *template.Template
	filename string
	Views    []View
}

func NewViewList(filename string) *ViewList {
	t := template.Must(template.New("viewListTemplate").Parse(viewListTemplate))

	return &ViewList{
		template: t,
		filename: filename,
		Views:    make([]View, 0),
	}
}

func (vl *ViewList) String() string {
	var buf bytes.Buffer
	err := vl.template.Execute(&buf, vl)
	if err != nil {
		panic(err)
	}
	return buf.String()
}

func (vl *ViewList) AddView(view View) {
	vl.Views = append(vl.Views, view)
}
