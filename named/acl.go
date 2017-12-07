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
	Elements map[string]CidrAddress
}

func NewAcl(name string) *Acl {
	t := template.Must(template.New("aclTemplate").Parse(aclTemplate))

	return &Acl{
		Name:     name,
		template: t,
		Elements: make(map[string]CidrAddress, 0),
	}
}

func (a Acl) String() string {
	var buf bytes.Buffer
	err := a.template.Execute(&buf, a)
	if err != nil {
		panic(err)
	}
	return buf.String()
}

func (a *Acl) Add(key string, element CidrAddress) {
	a.Elements[key] = element
}

func (a *Acl) Delete(key string) {
	delete(a.Elements, key)
}

func (a Acl) Contains(key string) bool {
	for k := range a.Elements {
		if k == key {
			return true
		}
	}
	return false
}

func (a Acl) Len() int {
	return len(a.Elements)
}
