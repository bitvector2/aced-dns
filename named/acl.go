package named

import (
	"bytes"
	"text/template"
)

const (
	aclTemplate = `
acl {{ .Name }} {
    {{- range .Elements }}
    {{ . }};
    {{- end }}
};
`
)

type Acl struct {
	Name     string
	template *template.Template
	Elements []CidrAddress
}

func NewAcl(name string) *Acl {
	t := template.Must(template.New("aclTemplate").Parse(aclTemplate))

	return &Acl{
		Name:     name,
		template: t,
		Elements: make([]CidrAddress, 0),
	}
}

func (a *Acl) String() string {
	var buf bytes.Buffer
	err := a.template.Execute(&buf, a)
	if err != nil {
		panic(err)
	}
	return buf.String()
}

func (a *Acl) AddElement(element CidrAddress) {
	a.Elements = append(a.Elements, element)
}
