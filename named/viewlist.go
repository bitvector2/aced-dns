package named

import (
	"bytes"
	"fmt"
	"os"
	"text/template"

	"github.com/bitvector2/testgo/utils"
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

func NewViewList(outputDir string) *ViewList {
	t := template.Must(template.New("viewListTemplate").Parse(viewListTemplate))

	return &ViewList{
		template: t,
		filename: fmt.Sprintf("%s/named.conf.viewlist", outputDir),
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

func (vl *ViewList) Save() {
	for v := 0; v < len(vl.Views); v++ {
		vl.Views[v].Save()
	}

	var buf bytes.Buffer
	buf.WriteString(vl.String())
	utils.CreateFile(vl.filename, buf.Bytes(), os.FileMode(0666))
}
