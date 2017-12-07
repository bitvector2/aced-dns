package named

import (
	"bytes"
	"net"
	"text/template"
)

const (
	viewTemplate = `
view {{ .Name }} {
    match-clients {
        {{ .Clients.Name }};
    };

    recursion no;

    forwarders {
        {{- range .Forwarders }}
        {{ . }};
        {{- end }}
    };
};
`
)

type View struct {
	Name       string
	template   *template.Template
	Clients    Acl
	Forwarders map[string]net.IP
}

func NewView(name string) *View {
	t := template.Must(template.New("viewTemplate").Parse(viewTemplate))

	return &View{
		Name:       name,
		template:   t,
		Clients:    *NewAcl(name), // aclname and viewname must match
		Forwarders: make(map[string]net.IP, 0),
	}
}

func (v View) String() string {
	var buf bytes.Buffer
	err := v.template.Execute(&buf, v)
	if err != nil {
		panic(err)
	}
	return buf.String()
}

func (v *View) Add(ip net.IP) {
	v.Forwarders[ip.String()] = ip
}

func (v *View) Delete(ip net.IP) {
	delete(v.Forwarders, ip.String())
}

func (v View) Contains(ip net.IP) bool {
	for _, v := range v.Forwarders {
		if v.String() == ip.String() {
			return true
		}
	}
	return false
}

func (v View) Len() int {
	return len(v.Forwarders)
}
