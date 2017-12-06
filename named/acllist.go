package named

import (
	"bytes"
	"text/template"
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

func NewAclList(filename string) *AclList {
	t := template.Must(template.New("aclListTemplate").Parse(aclListTemplate))

	return &AclList{
		template: t,
		filename: filename,
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
