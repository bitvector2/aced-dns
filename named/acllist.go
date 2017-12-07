package named

import (
	"bytes"
	"fmt"
	"os"
	"text/template"

	"github.com/bitvector2/testgo/utils"
)

const (
	aclListTemplate = `
    {{- range .Acls }}
    {{ . }}
    {{- end }}
`
)

type AclList struct {
	template *template.Template
	filename string
	Acls     []Acl
}

func NewAclList(outputDir string) *AclList {
	t := template.Must(template.New("aclListTemplate").Parse(aclListTemplate))

	return &AclList{
		template: t,
		filename: fmt.Sprintf("%s/named.conf.acllist", outputDir),
		Acls:     make([]Acl, 0),
	}
}

func (al *AclList) String() string {
	var buf bytes.Buffer
	err := al.template.Execute(&buf, al)
	if err != nil {
		panic(err)
	}
	return buf.String()
}

func (al *AclList) AddAcl(acl Acl) {
	al.Acls = append(al.Acls, acl)
}

func (al *AclList) Save() {
	var buf bytes.Buffer
	buf.WriteString(al.String())
	utils.CreateFile(al.filename, buf.Bytes(), os.FileMode(0666))
}
