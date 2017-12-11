package named

import (
	"bytes"
	"fmt"
	"net"
	"os"
	"text/template"

	"github.com/bitvector2/aced-dns/utils"
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
	Views    map[string]View
}

func NewViewList(outputDir string) *ViewList {
	t := template.Must(template.New("viewListTemplate").Parse(viewListTemplate))

	return &ViewList{
		template: t,
		filename: fmt.Sprintf("%s/named.conf.viewlist", outputDir),
		Views:    make(map[string]View, 0),
	}
}

func (vl ViewList) String() string {
	var buf bytes.Buffer
	err := vl.template.Execute(&buf, vl)
	if err != nil {
		panic(err)
	}
	return buf.String()
}

func (vl *ViewList) Add(viewname string) {
	vl.Views[viewname] = *NewView(viewname)
}

func (vl *ViewList) Delete(viewname string) {
	delete(vl.Views, viewname)
}

func (vl ViewList) Save() bool {
	var buf bytes.Buffer
	buf.WriteString(vl.String())
	dirty, err := utils.UpdateFile(vl.filename, buf.Bytes(), os.FileMode(0666))
	utils.Check(err)
	return dirty
}

func (vl ViewList) Contains(viewname string) bool {
	for k, _ := range vl.Views {
		if k == viewname {
			return true
		}
	}
	return false
}

func (vl ViewList) AddForwarder(viewname string, ip string) {
	obj := vl.Views[viewname]
	obj.Add(net.ParseIP(ip))
}

func (vl ViewList) DelForwarder(ip net.IP) {
	for _, v := range vl.Views {
		if v.Contains(ip) {
			v.Delete(ip)
		}
	}
}
